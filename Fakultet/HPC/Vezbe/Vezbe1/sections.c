#include <stdio.h>
#include <omp.h>

int main() {
#pragma omp parallel 
{
    #pragma omp sections
    {
        printf("Section 1.\n");

        #pragma omp section
        {
            printf("Section 2.\n");
        }

        #pragma omp section
        {
            printf("Section 3.\n");
        }

        #pragma omp section
        {
            printf("Section 4.\n");
        }

        #pragma omp section
        {
            printf("Section 5.\n");
        }

        #pragma omp section
        {
            printf("Section 6.\n");
        }

        #pragma omp section
        {
            printf("Section 7.\n");
        }


        #pragma omp section
        {
            printf("Section 8.\n");
        }

        #pragma omp section
        {
            printf("Section 9.\n");
        }

        #pragma omp section
        {
            printf("Section 10.\n");
        }
    }
}

return 0;
}