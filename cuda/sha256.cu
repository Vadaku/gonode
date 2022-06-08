// cd /home/hork/cuda-workspace/CudaSHA256/Debug/files
// time ~/Dropbox/FIIT/APS/Projekt/CpuSHA256/a.out -f ../file-list
// time ../CudaSHA256 -f ../file-list


#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <cuda.h>
#include "sha256.cuh"
#include <dirent.h>
#include <ctype.h>

extern "C" void sha256_block_data_order (uint32_t *ctx, const void *in, size_t num);

char * trim(char *str){
    size_t len = 0;
    char *frontp = str;
    char *endp = NULL;

    if( str == NULL ) { return NULL; }
    if( str[0] == '\0' ) { return str; }

    len = strlen(str);
    endp = str + len;

    /* Move the front and back pointers to address the first non-whitespace
     * characters from each end.
     */
    while( isspace((unsigned char) *frontp) ) { ++frontp; }
    if( endp != frontp )
    {
        while( isspace((unsigned char) *(--endp)) && endp != frontp ) {}
    }

    if( str + len - 1 != endp )
            *(endp + 1) = '\0';
    else if( frontp != str &&  endp == frontp )
            *str = '\0';

    /* Shift the string so that it starts at str so that if it's dynamically
     * allocated, we can still free it on the returned pointer.  Note the reuse
     * of endp to mean the front of the string buffer now.
     */
    endp = str;
    if( frontp != str )
    {
            while( *frontp ) { *endp++ = *frontp++; }
            *endp = '\0';
    }


    return str;
}

__global__ void sha256_hash(JOB ** jobs, int n) {
	int i = blockIdx.x * blockDim.x + threadIdx.x;
	// perform sha256 calculation here
	if (i < n){
		SHA256_CTX ctx;
		sha256_init(&ctx);
		sha256_update(&ctx, jobs[i]->data, jobs[i]->size, i* jobs[i]->size);
		sha256_final(&ctx, jobs[i]->digest, i * 32);
	}
}

__global__ void sha256_cuda(BYTE * data, BYTE * digest, int n, int messageSize, BYTE * target,
        int targetHexCharCount, int targetLength, int * position, BYTE * d_data_init,
        uint64_t nonce, bool * found, uint64_t * nonces) {
    int i = blockIdx.x * blockDim.x + threadIdx.x;
    // perform sha256 calculation here
    if (i < n) {
        SHA256_CTX ctx;
        while (*found == false) {
            int buffer = i * messageSize;
            uint64_t tmp = nonce + i;

            // initialize nonce
            #pragma unroll
            for (int j = 1; j <= 20; j++) {
                data[buffer + messageSize - j] = (tmp % 10) + '0';
                tmp /= 10;
            }

            // initialize message data
            for (int j = 0; j < messageSize - 20; j++) {
                data[j + buffer] = d_data_init[j];
            }

            // hash data
            sha256_init(&ctx);
            sha256_update(&ctx, data, messageSize, i * messageSize);
            sha256_final(&ctx, digest, i * 32);

            position[i] = 1;

            // Check that the hash that is generated hash a valid target
            for (int j = 0; j < targetHexCharCount; j++) {
                unsigned int value = (unsigned int) digest[(i * 32) + j];
                // Check used if targetlength is an odd integer value
                if (targetLength % 2 != 0 && j == targetHexCharCount - 1) {
                    // Bitwise operation to check first value in hex val
                    if (value >> 4 != target[j]) {
                        position[i] = 0;
                        break;
                    }
                } else if (value != target[j]) {
                    // Check if hex values are not equal
                    position[i] = 0;
                    break;
                }
            }

            if (position[i] == 1) {
                nonces[i] = nonce + i;
                *found = true;
                return;
            }
            nonce += n;
        }
    }
}

void pre_sha256() {
	// compy symbols
	checkCudaErrors(cudaMemcpyToSymbol(dev_k, host_k, sizeof(host_k), 0, cudaMemcpyHostToDevice));
}

extern "C" {


void runJobs(JOB ** jobs, int n){
	int blockSize = 4;
	int numBlocks = (n + blockSize - 1) / blockSize;
	sha256_hash <<< numBlocks, blockSize >>> (jobs, n);
}

}


JOB * JOB_init(BYTE * data, long size, char * fname) {
	JOB * j;
	checkCudaErrors(cudaMallocManaged(&j, sizeof(JOB)));	//j = (JOB *)malloc(sizeof(JOB));
	checkCudaErrors(cudaMallocManaged(&(j->data), size));
	j->data = data;
	j->size = size;
	for (int i = 0; i < 64; i++)
	{
		j->digest[i] = 0xff;
	}
	strcpy(j->fname, fname);
	return j;
}


BYTE * get_file_data(char * fname, unsigned long * size) {
	FILE * f = 0;
	BYTE * buffer = 0;
	unsigned long fsize = 0;

	f = fopen(fname, "rb");
	if (!f){
		fprintf(stderr, "get_file_data Unable to open '%s'\n", fname);
		return 0;
	}
	fflush(f);

	if (fseek(f, 0, SEEK_END)){
		fprintf(stderr, "Unable to fseek %s\n", fname);
		return 0;
	}
	fflush(f);
	fsize = ftell(f);
	rewind(f);

	//buffer = (char *)malloc((fsize+1)*sizeof(char));
	checkCudaErrors(cudaMallocManaged(&buffer, (fsize+1)*sizeof(char)));
	fread(buffer, fsize, 1, f);
	fclose(f);
	*size = fsize;
	return buffer;
}

BYTE * get_data(char* name, unsigned long * size){
	BYTE* buffer = 0;
	unsigned long ssize = 0;
	ssize = strlen(name);
	checkCudaErrors(cudaMallocManaged(&buffer, (ssize+1)*sizeof(char)));
	memcpy(buffer, name, ssize+1);
	*size = ssize;
	return buffer;

}

void print_usage(){
	printf("Usage: CudaSHA256 [OPTION] [FILE]...\n");
	printf("Calculate sha256 hash of given FILEs\n\n");
	printf("OPTIONS:\n");
	printf("\t-f FILE1 \tRead a list of files (separeted by \\n) from FILE1, output hash for each file\n");
	printf("\t-h       \tPrint this help\n");
	printf("\nIf no OPTIONS are supplied, then program reads the content of FILEs and outputs hash for each FILEs \n");
	printf("\nOutput format:\n");
	printf("Hash following by two spaces following by file name (same as sha256sum).\n");
	printf("\nNotes:\n");
	printf("Calculations are performed on GPU, each seperate file is hashed in its own thread\n");
}