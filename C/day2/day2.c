#include <stdio.h>
#include <stdlib.h>
#include <string.h> 

int hand_score(char c)
{
    switch (c) {
        case 'X':
            return 1;
        case 'Y':
            return 2;
        case 'Z':
            return 3;
    }
    return -1;
}

int match_socore(char op, char you)
{
    
    if ((op == 'A' && you == 'X') || (op == 'B' && you == 'Y') || (op == 'C' && you == 'Z')) return 3;
    if ((op == 'A' && you == 'Z') || (op == 'B' && you == 'X')|| (op == 'C' && you == 'Y')) return 0; 
    if ((op == 'A' && you == 'Y') || (op == 'B' && you == 'Z') || (op == 'C' && you == 'X')) return 6;
    return -1;
}

void solve_1(const char *file_name)
{

    FILE * fp; 
    
    fp = fopen(file_name, "r");
    if (fp == NULL)
        exit(EXIT_FAILURE);
    char chunk[128];
    printf("starting...\n");

    int score = 0; 

    while(fgets(chunk, sizeof(chunk), fp) != NULL) 
    {
        char a[2];
        char b[2];
        sscanf(chunk, "%s %s", a, b);
        int hand = hand_score(b[0]);
        if (hand == -1) exit(1);
        int match = match_socore(a[0], b[0]);
        if (match == -1) exit(1);

        score += hand + match;
    }

    fclose(fp);

    printf("solution for task 2.1 for data from %s is: %d\n", file_name, score);
}

int main() 
{

    char const *FILE_TASK = "./data_task.txt";

    solve_1(FILE_TASK);
    return 0;
}
