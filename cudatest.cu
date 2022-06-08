#include <stdio.h>
#include <cuda.h>
#include "cuda/sha256.cuh"
#include "cuda/sha256.cu"
 
//Executed on GPU

extern "C" {

    char* getGPU(char* stringy) {
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
}