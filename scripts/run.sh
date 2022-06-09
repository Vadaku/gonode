#!/bin/bash

echo "Compiling and running node"

nvcc --ptxas-options=-v --compiler-options '-fPIC' -o ../lib/libcudahash.so --shared ../cudatest.cu

cd ..

go run . 