#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stdbool.h>

const int CHUNK_SIZE = 256000;

bool has_occurance(char *c, int num)
{
    unsigned char bits[32];
    
    for(int i = 0; i < 32; i++ )
    {    
        bits[i] = 0;
    }

    for (int i = 0; i < num; i++)
    {
        int j = c[i]/8;
        unsigned char mask = 1<<(c[i]%8);
        if ((bits[j]&mask) != 0)
        {
            return true;
        }
        bits[j] |= mask;
    }

   return false;
}

int calculate_occurenece(char *c, int num)
{
    for (int i = 0; i < CHUNK_SIZE-num; i++)
    {
        if (c[i+num] == '\n' || c[i+num] == ' ') {
            return -1;
        }
        if (has_occurance(c+i, num)==false)
        {
            return i+num;
        }
    }

    return -1;
}

void solve(const char *file_name, int num)
{

    FILE * fp;

    fp = fopen(file_name, "r");
    if (fp == NULL) exit(EXIT_FAILURE);
    char chunk[CHUNK_SIZE];

    int result;

    while(fgets(chunk, sizeof(chunk), fp) != NULL)
    {
        if (strcmp(chunk, "\n") == 0)
        {
            continue;
        }
        result = calculate_occurenece(chunk, num);

    }
    printf("pile arrangement to task 5 slice size %d for data from %s is: %d \n",num, file_name, result);

}

int main()
{
    const char *TASK_FILE_NAME = "data_task.txt";
    solve(TASK_FILE_NAME, 4);
    solve(TASK_FILE_NAME, 14);
}