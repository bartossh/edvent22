#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <ctype.h>
#include <stdbool.h>

const int CHUNK_SIZE = 256;

char find_diff_char(char *a, int last)
{
    unsigned char bits[32];
    
    for(int i = 0; i < 32; i++ )
    {    
        bits[i] = 0;
    }

    for (int i = 0; i < last /2; i++)
    {
        int j = a[i]/8;
        unsigned char mask = 1<<(a[i]%8);
        bits[j] |= mask;
    }

    for (int i = last/2; i < last; i++)
    { 
        int j = a[i]/8;
        unsigned char mask = 1<<(a[i]%8);
        if ((bits[j]&mask) != 0)
        {
            return a[i]; 
        }
    }

    return '0';
}

char *get_repetitions(char *str_a,char *str_b)
{
    int i = 0; 
    unsigned char bits[32];
    
    for(int i = 0; i < 32; i++ )
    {    
        bits[i] = 0;
    }

    while(true)
    {
        char c = str_a[i];
        if (c == '\n')
        {
            break;
        }
        int j = c/8;
        unsigned char mask = 1<<(c%8);
        bits[j] |= mask;
        i++;
    }

    
    i = 0;
    int next = 0;
    char *repetions = calloc(CHUNK_SIZE, sizeof(char));
    while(true)
    {
        char c = str_b[i];
        if (c == '\n')
        {
            break;
        }
        int j = c/8;
        unsigned char mask = 1<<(c%8);
        if ((bits[j]&mask) != 0)
        {
            repetions[next] = c; 
            next++;
        }
        i++;
    }

    repetions[next+1] = '\n';
    return repetions;

}

int read_priority(char c)
{
    if (isupper(c) == 1)
    {
        return c - 65 + 27;
    }
    return c - 96;
}

int calc_priority(char s[CHUNK_SIZE])
{
    char *e;
    int index;
    e = strchr(s, '\n');
    index = (int)(e - s);
    char c = find_diff_char(s, index);
    return read_priority(c);
}

void solve_1(const char *file_name)
{

    FILE * fp;

    fp = fopen(file_name, "r");
    if (fp == NULL) exit(EXIT_FAILURE);
    char chunk[CHUNK_SIZE];

    int score = 0;

    while(fgets(chunk, sizeof(chunk), fp) != NULL)
    {
        if (strcmp(chunk, "\n") == 0)
        {
            continue;
        }
        score += calc_priority(chunk);
    }

    printf("solution to task 3.1 for data from %s is: %d\n", file_name, score);
}

void solve_2(const char *file_name)
{

    FILE * fp;

    fp = fopen(file_name, "r");
    if (fp == NULL) exit(EXIT_FAILURE);
    char chunk[CHUNK_SIZE];
    int score = 0;
    int counter = 0;
    char group[3][CHUNK_SIZE];
    while(fgets(chunk, sizeof(chunk), fp) != NULL)
    {
        if (strcmp(chunk, "\n") == 0)
        {
            continue;
        }

        memcpy(group[counter], chunk, sizeof(group[counter]));

        if (counter != 0 && counter %2 == 0)
        {
            char *repetions = get_repetitions(group[0], group[1]);
            char *final_rep = get_repetitions(repetions, group[2]);
            score += read_priority(final_rep[0]);
            counter = 0;
            continue;
        }
        counter++;
    }

    printf("solution to task 3.2 for data from %s is: %d\n", file_name, score);
}


int main()
{
    char const *FILE_TEST = "./data_test.txt"; 
    solve_1(FILE_TEST);
    solve_2(FILE_TEST);

    char const *FILE_TASK = "./data_task.txt"; 
    solve_1(FILE_TASK);
    solve_2(FILE_TASK);
    
    return 0;
}
