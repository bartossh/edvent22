#include <stdio.h>
#include <stdlib.h>
#include <string.h> 

int hand_score(char c)
{
    switch (c) 
    {
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

char wining_hand(char op)
{
    switch (op)
    {
        case 'A':
            return 'Y';
        case 'B':
            return 'Z';
        case 'C':
            return 'X';
    }

    return 'e';
}

char draw_hand(char op)
{
    switch (op)
    {
        case 'A':
            return 'X';
        case 'B':
            return 'Y';
        case 'C':
            return 'Z';
    }

    return 'e';
}

char loosing_hand(char op)
{
    switch (op)
    {
        case 'A':
            return 'Z';
        case 'B':
            return 'X';
        case 'C':
            return 'Y';
    }

    return 'e';
}

char transform_to_your_hand(char op, char act)
{
    switch (act)
    {
        case 'X':
            return loosing_hand(op);
        case 'Y':
            return draw_hand(op);
        case 'Z':
            return wining_hand(op);
    }

    return 'e'; // e - no coresponding match in other funcs than will exit with
}

void solve_1(const char *file_name)
{

    FILE * fp; 
    
    fp = fopen(file_name, "r");
    if (fp == NULL) exit(EXIT_FAILURE);
    char chunk[128];
    printf("starting...\n");

    int score = 0; 

    while(fgets(chunk, sizeof(chunk), fp) != NULL) 
    {
        char a[2];
        char b[2];
        sscanf(chunk, "%s %s", a, b);
        int hand = hand_score(b[0]);
        if (hand == -1) exit(EXIT_FAILURE);
        int match = match_socore(a[0], b[0]);
        if (match == -1) exit(EXIT_FAILURE);

        score += hand + match;
    }

    fclose(fp);

    printf("solution for task 2.1 for data from %s is: %d\n", file_name, score);
}

void solve_2(const char *file_name)
{

    FILE * fp; 
    
    fp = fopen(file_name, "r");
    if (fp == NULL) exit(EXIT_FAILURE);
    char chunk[128];
    printf("starting...\n");

    int score = 0; 

    while(fgets(chunk, sizeof(chunk), fp) != NULL) 
    {
        char a[2];
        char b[2];
        sscanf(chunk, "%s %s", a, b);
        char c = transform_to_your_hand(a[0], b[0]);
        int hand = hand_score(c);
        if (hand == -1) exit(EXIT_FAILURE);
        int match = match_socore(a[0], c);
        if (match == -1) exit(EXIT_FAILURE);

        score += hand + match;
    }

    fclose(fp);

    printf("solution for task 2.1 for data from %s is: %d\n", file_name, score);
}

int main() 
{

    char const *FILE_TASK = "./data_task.txt";

    solve_1(FILE_TASK);
    solve_2(FILE_TASK);
    return 0;
}
