#!/bin/bash

echo "Compiling and running node"

nvcc --ptxas-options=-v --compiler-options '-fPIC' -o ../lib/libgetgpu.so --shared ../cudatest.cu

cd ..

go run . 