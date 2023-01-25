#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stdbool.h>

const int CHUNK_SIZE = 256;
const int INITIAL_SIZE = 10;

typedef struct text_buffer {
    char **buffer;
    int len;
    int cap; 
} text_buffer;

void append_text_buffer(text_buffer *buf, char chunk[256])
{
    if (buf->len == buf->cap)
    {
        buf->cap += buf->cap;
        buf->buffer = realloc(buf->buffer, buf->cap*sizeof(char*));
        for (int i = buf->len; i < buf->cap; i++) buf->buffer[i] = malloc(CHUNK_SIZE*sizeof(char));
    }
    memcpy(buf->buffer[buf->len], chunk, sizeof(char)*CHUNK_SIZE);
    buf->len++;
}

text_buffer new_text_buffer()
{
    text_buffer b;
    b.buffer = malloc(INITIAL_SIZE*sizeof(char*));
    for (int i = 0; i < INITIAL_SIZE; i++) b.buffer[i] = malloc(CHUNK_SIZE*sizeof(char));
    b.cap = INITIAL_SIZE;
    b.len = 0;
    return b;
}

void iter_text_buffer(text_buffer *buf)
{
    for (int i = 0; i < buf->len; i++)
    {
        printf("%s\n", buf->buffer[i]);
    }
}

void free_text_buffer(text_buffer *buf)
{
    for (int i = 0; i < buf->cap; i++) free(buf->buffer[i]);
    free(buf->buffer); 
}

void solve_1(const char *file_name)
{

    FILE * fp;

    fp = fopen(file_name, "r");
    if (fp == NULL) exit(EXIT_FAILURE);
    char chunk[CHUNK_SIZE];

    int counter = 0;
    text_buffer text = new_text_buffer();


    while(fgets(chunk, sizeof(chunk), fp) != NULL)
    {
        if (strcmp(chunk, "\n") == 0)
            continue;
        
        append_text_buffer(&text, chunk);
        
    }

    iter_text_buffer(&text);

    printf("solution to task 5.1 for data from %s is: %d\n", file_name, counter);
    free_text_buffer(&text);
}

int main()
{
    const char *TEST_FILE_NAME = "data_test.txt";
    solve_1(TEST_FILE_NAME);

    const char *TASK_FILE_NAME = "data_task.txt";
    solve_1(TASK_FILE_NAME);
}