#include <stdio.h>
#include <stdlib.h>
#include <string.h>

const int INITIAL_ARR_SIZE = 10;

int compare( const void* a, const void* b)
{
     int int_a = * ( (int*) a );
     int int_b = * ( (int*) b );

     if ( int_a == int_b ) return 0;
     if ( int_a < int_b ) return -1;
     return 1;
}

typedef struct items{
    int len;
    int cap;
    int *values;
} items;

items new_items()
{
    items it; 
    it.values = malloc(sizeof(int)*INITIAL_ARR_SIZE);
    it.cap = INITIAL_ARR_SIZE;
    it.len = 0;
    return it;
}

void print_items_values(items it)
{
    for (int i = 0; i < it.len; i++) 
    {
        printf("%d\n", it.values[i]);
    }
}

void append_to_items(items *is, int i)
{
    if (is->len == is->cap) 
    {
        is->cap += is->cap; // grow buffer by the factor or times two the capacity
        is->values = realloc(is->values, is->cap*sizeof(int));
    }
    is->values[is->len] = i;
    is->len++;
}

int calc_calories_items(items it)
{
    int sum = 0;
    for (int i = 0; i < it.len; i++)
    {
        sum += it.values[i];
    }
    return sum;
} 

void free_items_values(items it) {
        free(it.values);
}

typedef struct elves {
    int len;
    int cap;
    items *values;
} elves;

elves new_elves()
{
    elves el; 
    el.values = malloc(sizeof(items)*INITIAL_ARR_SIZE);
    el.cap = INITIAL_ARR_SIZE;
    el.len = 0;
    return el;
}

void print_elves_values(elves el)
{
    for (int i= 0; i < el.len; i++) 
    {
        printf("\nnext:\n");
        print_items_values(el.values[i]);
    } 
}


void append_to_elves(elves *el, items i)
{
    if (el->len == el->cap) 
    {
        el->cap += el->cap; // grow buffer by the factor or times two the capacity
        el->values = realloc(el->values, el->cap*sizeof(items));
    }

    el->values[el->len] = i;
    el->len++;
}

int calc_max_cal_elves(elves el)
{
    int max = 0;
    for (int i = 0; i < el.len; i++)
    {
        int val = calc_calories_items(el.values[i]);
        if (val > max) 
        {
            max = val;
        }
    }
    return max;
}

int calc_max_cal_top_elves(elves el, int top)
{
    if (top > el.len) 
    {
        return -1; // avoid accessing values from outside of a buffer
    }
    
    int *calories = malloc(sizeof(int)*el.len);

    for (int i = 0; i < el.len; i++)
    {
        int val = calc_calories_items(el.values[i]);
        calories[i] = val;
    }
    qsort( calories, el.len, sizeof(int), compare);
    
    int sum = 0;
    for (int i = el.len-1; i > el.len-1-top; i--) 
    {
        sum += calories[i];
    } 
    
    free(calories);

    return sum;
}

void free_elves_values(elves el)
{
    for (int i = 0; i < el.len; i++) 
    {
        free_items_values(el.values[i]);
    }
    free(el.values);
}

void solve_1(const char *file_name)
{

    FILE * fp; 
    
    fp = fopen(file_name, "r");
    if (fp == NULL)
        exit(EXIT_FAILURE);
    char chunk[128];
    
    elves el = new_elves();
    items it = new_items();

    while(fgets(chunk, sizeof(chunk), fp) != NULL) 
    {
        if (strcmp(chunk, "\n") == 0)
        {
            append_to_elves(&el, it);
            it = new_items();
            continue;
        }
        int x = atoi(chunk);
        append_to_items(&it, x); 
    } 
    fclose(fp);

    printf("size of int is %lu\n", sizeof(int));
    int res = calc_max_cal_elves(el);
    printf("max clalories task 1.1 for el max claories for data from %s is: %d\n", file_name, res);
    free_elves_values(el);
}

void solve_2(const char *file_name)
{

    FILE * fp; 
    
    fp = fopen(file_name, "r");
    if (fp == NULL)
        exit(EXIT_FAILURE);
    char chunk[128];
    
    elves el = new_elves();
    items it = new_items();

    while(fgets(chunk, sizeof(chunk), fp) != NULL) 
    {
        if (strcmp(chunk, "\n") == 0)
        {
            append_to_elves(&el, it);
            it = new_items();
            continue;
        }
        int x = atoi(chunk);
        append_to_items(&it, x); 
    } 
    fclose(fp);

    printf("size of int is %lu\n", sizeof(int));
    int res = calc_max_cal_top_elves(el, 3);
    printf("max clalories task 1.2 for top 3 elves for data from %s is: %d\n", file_name, res);
    free_elves_values(el);
}

int main() 
{

    char const *FILE_TEST = "./data_test.txt"; 
    char const *FILE_TASK = "./data_task.txt";

    solve_1(FILE_TEST);
    solve_1(FILE_TASK);
    solve_2(FILE_TEST);
    solve_2(FILE_TASK);
}
