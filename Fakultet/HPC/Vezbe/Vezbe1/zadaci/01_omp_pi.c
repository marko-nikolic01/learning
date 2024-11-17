#include <stdio.h>
#include <omp.h>

#define NUM_THREADS 2
#define PADDING 8

#define PRINT_SEQ_RES 1
#define PRINT_PRL_RES_BASIC 1
#define PRINT_PRL_RES_NO_FALSE_SHARING 1
#define PRINT_PRL_RES_SYNCHRONIZATION 1
#define PRINT_PRL_RES_FOR_CONSTRUCT 1

static long num_steps = 100000;

double step;
void serial_code();
 void parallel_code_incorrect();
void parallel_code();
void parallel_code_no_false_sharing();
void parallel_code_synchronization();
void parallel_code_for_construct();

int main() {

    printf("*************** Sekvencijalna implementacija ***************\n");
    serial_code();
    printf("************************************************************\n");

    printf("************ Neispravna paralelna implementacija ************\n");
    parallel_code_incorrect();
    printf("************************************************************\n");

    printf("***************** Paralelna implementacija *****************\n");
    parallel_code();
    printf("************************************************************\n");

    printf("***** Paralelna implementacija (nema laznog deljenja) ******\n");
    parallel_code_no_false_sharing();
    printf("************************************************************\n");

    printf("******** Paralelna implementacija (sinhronizacija) *********\n");
    parallel_code_synchronization();
    printf("************************************************************\n");

    printf("******** Paralelna implementacija (for konstrukcija) ********\n");
    parallel_code_for_construct();
    printf("************************************************************\n");

    return 0;
}

void serial_code() {
    double start = omp_get_wtime();
    int i;
    double x, pi, sum = 0.0;

    step = 1.0 / (double) num_steps;

    for (i = 0; i < num_steps; i++) { 
        double x = (i + 0.5) * step;
        sum = sum + 4.0 / (1.0 + x * x);
    }

    pi = step * sum;
    double end = omp_get_wtime();

    printf("pi = %lf\n", pi);
    printf("Time elapsed: %lf\n", end - start);
}

void parallel_code_incorrect() {
    double start = omp_get_wtime();
    int i;
    double x, pi, sum = 0.0;

    step = 1.0 / (double) num_steps;

    #pragma omp parallel
    {
        for (i = 0; i < num_steps; i++) { 
            double x = (i + 0.5) * step;
            sum = sum + 4.0 / (1.0 + x * x);
        }
    }

    pi = step * sum;
    double end = omp_get_wtime();

    printf("pi = %lf\n", pi);
    printf("Time elapsed: %lf\n", end - start);
}

void parallel_code() {
    double start = omp_get_wtime();
    int i, nthreads;
    double x, pi, sum[NUM_THREADS];

    step = 1.0 / (double) num_steps;

    omp_set_num_threads(NUM_THREADS);
    #pragma omp parallel
    {
        int id = omp_get_thread_num();
        int n = omp_get_num_threads();

        sum[id] = 0.0;

        if(id == 0) {
            nthreads = n;
        }

        for (int i = id; i < num_steps; i += n) { 
            double x = (i + 0.5) * step;
            sum[id] += 4.0 / (1.0 + x * x);
        }
    }

    for (int i = 0; i < nthreads; i++) { 
        pi += sum[i];
    }

    pi *= step;
    double end = omp_get_wtime();

    printf("pi = %lf\n", pi);
    printf("Time elapsed: %lf\n", end - start);
}

void parallel_code_no_false_sharing() {
    double start = omp_get_wtime();
    int i, nthreads;
    double x, pi, sum[NUM_THREADS][PADDING];

    step = 1.0 / (double) num_steps;

    omp_set_num_threads(NUM_THREADS);
    #pragma omp parallel
    {
        int id = omp_get_thread_num();
        int n = omp_get_num_threads();

        sum[id][0] = 0.0;

        if(id == 0) {
            nthreads = n;
        }

        for (int i = id; i < num_steps; i += n) { 
            double x = (i + 0.5) * step;
            sum[id][0] += 4.0 / (1.0 + x * x);
        }
    }

    for (int i = 0; i < nthreads; i++) { 
        pi += sum[i][0];
    }

    pi *= step;
    double end = omp_get_wtime();

    printf("pi = %lf\n", pi);
    printf("Time elapsed: %lf\n", end - start);
}

void parallel_code_synchronization() {
    double start = omp_get_wtime();
    int i, nthreads;
    double x, pi;

    step = 1.0 / (double) num_steps;

    omp_set_num_threads(NUM_THREADS);
    #pragma omp parallel
    {
        int id = omp_get_thread_num();
        int n = omp_get_num_threads();

        double sum = 0.0;

        if(id == 0) {
            nthreads = n;
        }

        for (int i = id; i < num_steps; i += n) { 
            double x = (i + 0.5) * step;
            sum += 4.0 / (1.0 + x * x);
        }

        #pragma omp critical
        {
            pi += sum;
        }
    }

    pi *= step;
    double end = omp_get_wtime();

    printf("pi = %lf\n", pi);
    printf("Time elapsed: %lf\n", end - start);
}

void parallel_code_for_construct() {
    double start = omp_get_wtime();
    double pi = 0.0;

    step = 1.0 / (double) num_steps;

    omp_set_num_threads(NUM_THREADS);
    #pragma omp parallel for reduction(+:pi)
    for (int i = 0; i < num_steps; i++) { 
        double x = (i + 0.5) * step;
        pi += 4.0 / (1.0 + x * x);
    }

    pi *= step;
    double end = omp_get_wtime();

    printf("pi = %lf\n", pi);
    printf("Time elapsed: %lf\n", end - start);
}