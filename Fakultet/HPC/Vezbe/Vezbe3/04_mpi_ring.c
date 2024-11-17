#include <stdio.h>
#include <string.h>
#include <mpi.h>

#define MAX_STRIKES 10

int main(int argc, char* argv[]) {
    int size, rank;

    int token = -1;

    MPI_Init(&argc, &argv);

    MPI_Comm_size(MPI_COMM_WORLD, &size);
    MPI_Comm_rank(MPI_COMM_WORLD, &rank);

    --size;

    if(rank == 0) {
        MPI_Send(&token, 1, MPI_INT, 1, 0, MPI_COMM_WORLD);

        MPI_Recv(&token, 1, MPI_INT, size, 0, MPI_COMM_WORLD, MPI_STATUS_IGNORE);
        printf("Process %d received token %d from process %d.\n", rank, token, size);
    } else if(rank < size) {
        MPI_Recv(&token, 1, MPI_INT, rank - 1, 0, MPI_COMM_WORLD, MPI_STATUS_IGNORE);
        printf("Process %d received token %d from process %d.\n", rank, token, rank - 1);

        MPI_Send(&token, 1, MPI_INT, rank + 1, 0, MPI_COMM_WORLD);
    } else {
        MPI_Recv(&token, 1, MPI_INT, rank - 1, 0, MPI_COMM_WORLD, MPI_STATUS_IGNORE);
        printf("Process %d received token %d from process %d.\n", rank, token, rank - 1);

        MPI_Send(&token, 1, MPI_INT, 0, 0, MPI_COMM_WORLD);
    }

    MPI_Finalize();

    return 0;
}