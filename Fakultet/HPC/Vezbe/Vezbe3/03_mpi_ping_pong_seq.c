#include <stdio.h>
#include <string.h>
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

            char message[60];
            sprintf(message, "p%d sent ping_pong_count to p%d and incremented it to %d.\n", rank, opponent, ball);
            MPI_Send(message, strlen(message) + 1, MPI_CHAR, 2, ball, MPI_COMM_WORLD);
        } else if (rank == (ball + 1) % 2) {
            MPI_Recv(&ball, 1, MPI_INT, opponent, 0, MPI_COMM_WORLD, MPI_STATUS_IGNORE);
        } else {
            ball++;
            char message[60];
            MPI_Recv(message, 60, MPI_CHAR, (ball - 1) % 2, ball, MPI_COMM_WORLD, MPI_STATUS_IGNORE);
            printf("%s", message);
        }
    }

    MPI_Finalize();

    return 0;
}