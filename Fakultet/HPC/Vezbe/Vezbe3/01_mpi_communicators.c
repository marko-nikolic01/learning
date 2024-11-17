#include <stdio.h>
#include <mpi.h>

int main(int argc, char* argv[]) {
    int size, rank, newSize, newRank;

    MPI_Init(&argc, &argv);

    MPI_Comm_size(MPI_COMM_WORLD, &size);
    MPI_Comm_rank(MPI_COMM_WORLD, &rank);

    MPI_Comm newComm;

    MPI_Comm_split(MPI_COMM_WORLD, rank % 2, rank, &newComm);

    MPI_Comm_size(newComm, &newSize);
    MPI_Comm_rank(newComm, &newRank);

    printf("MPI_COMM_WORLD - (SIZE: %d, RANK: %d)\t newComm - (SIZE: %d, RANK: %d)\n", size, rank, newSize, newRank);

    MPI_Finalize();

    return 0;
}