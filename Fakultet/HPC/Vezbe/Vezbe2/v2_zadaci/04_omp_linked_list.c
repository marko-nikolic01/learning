#include <stdlib.h>
#include <stdio.h>
#include <omp.h>

#ifndef N
#define N 5
#endif
#ifndef FS
#define FS 38
#endif

struct node {
    int data;
    int fibdata;
    struct node* next;
};

int fib(int n) {
    int x, y;
    if (n < 2) {
        return (n);
    } else {
        x = fib(n - 1);
        y = fib(n - 2);
        return (x + y);
    }
}

// Funkcija koja procesira svaki element liste. Za svaki element
// liste se racuna vrednost n-tog elementa Fibonacijevog niza.
// Zadatak funkcije nije narocito bitan. Ono sto jeste bitno je da
// je racunanje n-tog elementa Fibonacijevog niza vremenski zahtevan
// posao.
void processwork(struct node* p) 
{
    int n;
    n = p->data;
    p->fibdata = fib(n);
}

// Inicijalizacija liste. U listu se uvezuje N elemenata i svakom
// se dodeljuje redni broj elementa Fibonacijevog niza koji je
// potrebno sracunati (od FS + 1, pa nadalje).
struct node* init_list(struct node* p) {
    int i;
    struct node* head = NULL;
    struct node* temp = NULL;
    
    head = malloc(sizeof(struct node));
    p = head;
    p->data = FS;
    p->fibdata = 0;
    for (i = 0; i < N; i++) {
        temp = malloc(sizeof(struct node));
        p->next = temp;
        p = temp;
        p->data = FS + i + 1;
        p->fibdata = i + 1;
    }
    p->next = NULL;
    return head;
}

int main(int argc, char *argv[]) {
    double start, end;
    struct node *p = NULL;
    struct node *temp = NULL;
    struct node *head = NULL;
    
	printf("Process linked list\n");
    printf("  Each linked list node will be processed by function 'processwork()'\n");
    printf("  Each ll node will compute %d fibonacci numbers beginning with %d\n", N, FS);      

    p = init_list(p);
    head = p;
    start = omp_get_wtime();

    #pragma omp parallel
    {  
        #pragma omp single
        {
            while (p != NULL) {
                { 
                    #pragma omp task firstprivate(p)
                    processwork(p);
                }

                p = p->next;
            }
        }
    }
    end = omp_get_wtime();

    // Oslobadjanje memorije zauzete za listu.
    p = head;
    while (p != NULL) {
        printf("%d : %d\n", p->data, p->fibdata);
        temp = p->next;
        free(p);
        p = temp;
    }  
    free(p);

    printf("Compute Time: %f seconds\n", end - start);
    return 0;
}

