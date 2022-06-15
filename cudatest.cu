#include <stdio.h>
#include <cuda.h>
#include "cuda/sha256.cuh"
#include "cuda/sha256.cu"
#include <string>
#include <vector>
#include <iostream>
 
//Executed on GPU

extern "C" {

    char* cudaHash(char* stringy) {
        JOB ** jobs;
        unsigned long temp;
        char * a_file = 0;
        BYTE * buff;
        int n = 0;

        a_file = stringy;

        buff = get_data(a_file, &temp);

        checkCudaErrors(cudaMallocManaged(&jobs, 1 * sizeof(JOB *)));
        jobs[n++] = JOB_init(buff, temp, a_file);
    
        pre_sha256();
        runJobs(jobs, n);
        cudaDeviceSynchronize();

        char * out;
        out = hash_to_string(jobs[0]->digest);

        // print_job(jobs[0]);

        return out;

        // cudaDeviceReset();
        // int nDevices;
        // cudaGetDeviceCount(&nDevices);
        // for (int i = 0; i < nDevices; i++) {
        //     cudaDeviceProp prop;
        //     cudaGetDeviceProperties(&prop, i);
        //     printf("Device Number: %d\n", i);
        //     printf("  Device name: %s\n", prop.name);
        //     printf("  Memory Clock Rate (KHz): %d\n",
        //         prop.memoryClockRate);
        //     printf("  Memory Bus Width (bits): %d\n",
        //         prop.memoryBusWidth);
        //     printf("  Peak Memory Bandwidth (GB/s): %f\n\n",
        //         2.0*prop.memoryClockRate*(prop.memoryBusWidth/8)/1.0e6);
        // }
    }

//     

    void cudaMine(char* s, char* t, uint64_t nonce,
    char* d, char* u, int numMessages, uint64_t timestamp) {

    std::string source(s);
    std::string target(t);
    std::string data(d);
    std::string user(u);
    
    std::string datahash = cudaHash(d);
    ++nonce;
    int targetLength = target.length() / 2;
    if (target.length() % 2 != 0) {
        targetLength++;
    }

    BYTE * d_digest;
    BYTE * d_data;
    BYTE * d_data_init;
    BYTE * d_target;
    int * d_position;
    uint64_t * d_nonces;

    const int nonceSize = 20;

    std::vector<std::string> resultHashes(numMessages, "");
    std::string src;
    if (source.length() == 64 && source.find_first_not_of("0123456789abcdefABCDEF") == std::string::npos) {
        // source is a hash
        src = source;
    } else {
        src = cudaHash(s);
    }
    std::string message = src + datahash + target + user + std::to_string(timestamp);
    int messageSize = message.length() + nonceSize;
    int dataSize = sizeof(BYTE) * numMessages * messageSize;

    // Allocate device variables in Unified Memory
    checkCudaErrors(cudaMallocManaged(&d_data_init, sizeof(BYTE) * message.length()));
    checkCudaErrors(cudaMallocManaged(&d_target, sizeof(BYTE) * targetLength));
    checkCudaErrors(cudaMallocManaged(&d_data, dataSize));
    checkCudaErrors(cudaMallocManaged(&d_digest, sizeof(BYTE) * numMessages * 32));
    checkCudaErrors(cudaMallocManaged(&d_position, sizeof(int) * numMessages));
    checkCudaErrors(cudaMallocManaged(&d_nonces, sizeof(uint64_t) * numMessages));

    checkCudaErrors(cudaMemcpy(d_data_init, &message.c_str()[0], sizeof(BYTE) * message.length(),
    cudaMemcpyHostToDevice));

    // Convert target to hex values and add to d_target host variable
    for (int i = 0; i < target.length(); i += 2) {
        // if last two elements in string get last two values
        std::string str;
        str.append(1, target[i]);
        if (i + 1 < target.length()) str.append(1, target[i + 1]);
        d_target[i / 2] = std::stoi(str, 0, 16);
    }

    int blockSize = 512;
    bool * found;

    checkCudaErrors(cudaMallocManaged(&found, sizeof(bool)));
    *found = false;

    int numBlocks = numMessages / blockSize;

    sha256_cuda<<<numBlocks, blockSize>>>(d_data, d_digest, numMessages, messageSize, d_target,
                targetLength, target.length(), d_position, d_data_init, nonce, found,
                d_nonces);
    checkCudaErrors(cudaDeviceSynchronize());

    // Find value in resulthashes that has a length greater than 0 and return it
    for (int i = 0; i < numMessages; i++) {
        if (d_position[i] == 1) {
            BYTE bdata[32];
            int start = i * 32;
            int end = start + 32;
            // Convert byte hash to string
            for (int j = start; j < end; j++) {
                bdata[j % 32] = d_digest[j];
            }

            std::string foundhash = hash_to_string(bdata);
            std::cout << foundhash << std::endl;
            break;
        }
    }

    checkCudaErrors(cudaFree(d_digest));
    checkCudaErrors(cudaFree(d_data));
    checkCudaErrors(cudaFree(d_target));
    checkCudaErrors(cudaFree(d_data_init));
    checkCudaErrors(cudaFree(d_position));
}
}