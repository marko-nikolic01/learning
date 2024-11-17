#include <stdio.h>
#include <mpi.h>

#define MAX_STRIKES 10

int main(int argc, char* argv[]) {
    int rank, ball = 0;

    MPI_Init(&argc, &argv);

    MPI_Comm_rank(MPI_COMM_WORLD, &rank);
    int opponent = (rank + 1) % 2;

    while(ball < MAX_STRIKES) {
        if(rank == ball % 2) {
            ball++;

            MPI_Send(&ball, 1, MPI_INT, opponent, 0, MPI_COMM_WORLD);

            printf("p%d sent ping_pong_count to p%d and incremented it to %d.\n", rank, opponent % 2, ball);
        } else {
            MPI_Recv(&ball, 1, MPI_INT, opponent, 0, MPI_COMM_WORLD, MPI_STATUS_IGNORE);

            printf("p%d received ping_pong_count %d from p%d.\n", rank, ball, opponent % 2);
        }
    }

    MPI_Finalize();

    return 0;
}