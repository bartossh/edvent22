#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stdbool.h>

const int CHUNK_SIZE = 128;

bool covers(int lx, int rx, int ly, int ry)
{

    if (lx <= ly && rx >= ry) return true;
    if (ly <= lx && ry >= rx) return true;
    
    return false;
}

bool overlaps(int lx, int rx, int ly, int ry)
{
    if ((rx >= ly && rx <= ry) || (ry >=lx && ry <=rx)) return true;
    return false;
}

void solve_1(const char *file_name)
{

    FILE * fp;

    fp = fopen(file_name, "r");
    if (fp == NULL) exit(EXIT_FAILURE);
    char chunk[CHUNK_SIZE];

    int counter = 0;

    while(fgets(chunk, sizeof(chunk), fp) != NULL)
    {
        if (strcmp(chunk, "\n") == 0)
        {
            continue;
        }
        int lx, rx, ly, ry;
        sscanf(chunk, "%d-%d,%d-%d,", &lx, &rx, &ly, &ry);
        if (covers(lx, rx, ly, ry))
        {
            counter++;
        }
    }

    printf("solution to task 4.1 for data from %s is: %d\n", file_name, counter);
}

void solve_2(const char *file_name)
{

    FILE * fp;

    fp = fopen(file_name, "r");
    if (fp == NULL) exit(EXIT_FAILURE);
    char chunk[CHUNK_SIZE];

    int counter = 0;

    while(fgets(chunk, sizeof(chunk), fp) != NULL)
    {
        if (strcmp(chunk, "\n") == 0)
        {
            continue;
        }
        int lx, rx, ly, ry;
        sscanf(chunk, "%d-%d,%d-%d,", &lx, &rx, &ly, &ry);
        if (overlaps(lx, rx, ly, ry))
        {
            counter++;
        }
    }

    printf("solution to task 4.2 for data from %s is: %d\n", file_name, counter);
}

int main()
{
    const char *TEST_FILE_NAME = "./data_test.txt";
    solve_1(TEST_FILE_NAME);
    solve_2(TEST_FILE_NAME);
    const char *TASK_FILE_NAME = "./data_task.txt";
    solve_1(TASK_FILE_NAME);
    solve_2(TASK_FILE_NAME);
}