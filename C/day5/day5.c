#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stdbool.h>

const int CHUNK_SIZE = 256;
const int INITIAL_SIZE = 10;
const int STACK_SIZE = 128;
const int STACKS_IN_WAREHOUSE = 12;

typedef struct pile {
    char *stack;
    int len;
    int cap;
} pile;

pile new_pile()
{
    pile p;
    p.len = 0;
    p.cap = STACK_SIZE;
    p.stack = calloc(STACK_SIZE, sizeof(char));
    return p;
}

int push_to_pile(pile *p, char c)
{
    if (p->len == p->cap)
    {
        return -1;
    }
    p->len++;
    for (int i = p->len-1; i > 0; i--)
    {
        p->stack[i] = p->stack[i-1];
    }
    p->stack[0] = c;
    return 0;
}

int append_to_pile(pile *p, char c)
{
    if (p->len == p->cap) {
        return -1;
    }
    p->stack[p->len] = c;
    p->len++;
    return p->len;
}

int append_many_to_pile(pile *p, char *c, int num)
{
    if (p->len + num > STACK_SIZE)
    {
        return -1;
    }
    for (int i = 0; i < num; i++)
    {
        p->stack[p->len] = c[i];
        p->len++;
    }
    return p->len;
}

char remove_from_pile(pile *p)
{
    if (p->len == 0) {
        return '\n';
    }
    p->len--;
    char c = p->stack[p->len];
    p->stack[p->len] = ' ';
    return c;
}

char* remove_many_from_pile(pile *p, int *num)
{

    if (p->len < (*num)) {
        *num = p->len;
    }
    char *c = malloc((*num)*sizeof(char));
    memcpy(c, p->stack+(p->len-(*num)), (*num));
    p->len -= (*num);
    return c;
}

char get_top_char_pile(pile *p)
{
    return p->stack[p->len-1];
}

void free_pile(pile *p)
{
    free(p->stack);
}

typedef struct warehouse {
    int cap;
    pile piles[STACKS_IN_WAREHOUSE];
} warehouse;

warehouse new_warehouse()
{
    warehouse w;
    w.cap = STACKS_IN_WAREHOUSE;
    for (int i = 0; i < w.cap; i++)
    {
        w.piles[i] = new_pile();
    }
    return w;
}

void read_line_to_warehouse(char s[CHUNK_SIZE], warehouse *w)
{
    for (int i = 0; i < CHUNK_SIZE; i++)
    {
        if (s[i] != ' ' && s[i] != '\n' && (i-1)/4 < w->cap)
        {
            if (i == 1 || (i - 1)%4==0)
            {   
                push_to_pile(&(w->piles[(i-1)/4]), s[i]);
            }
        }

    }
}

void print_warehouse(warehouse *w)
{
    for (int i = 0; i < w->cap; i++)
    {
        printf("stack: %i | %s \n", i, w->piles[i].stack);
    }
}

void free_warehouse(warehouse *w)
{
    for (int i = 0; i < STACKS_IN_WAREHOUSE; i++)
    {
        free_pile(&(w->piles[i]));
    }
}

char* get_top_chars_in_warehouse(warehouse *w, int num)
{
    char *c = malloc(num*sizeof(char));
    for (int i = 0; i < num; i++)
    {
        c[i] = get_top_char_pile(&(w->piles[i])); 
    }
    return c;
}

void single_move_piles_warehouse(warehouse *w, int m, int f, int t)
{
    for (int i = m; i > 0; i--)
    {
        char c = remove_from_pile(&(w->piles[f-1]));
        if (c == '\n')
        {
            break;
        }
        int idx = append_to_pile(&(w->piles[t-1]), c);
        if (idx < 0) exit(EXIT_FAILURE);
    }
}

void multiple_move_piles_warehouse(warehouse *w, int m, int f, int t)
{
    char *c = remove_many_from_pile(&(w->piles[f-1]), &m);
    int idx = append_many_to_pile(&(w->piles[t-1]), c, m);
    if (idx == -1) exit(EXIT_FAILURE);
    free(c);
}

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

void free_text_buffer(text_buffer *buf)
{
    for (int i = 0; i < buf->cap; i++) free(buf->buffer[i]);
    free(buf->buffer); 
}

void append_to_warehouse(text_buffer *buf, warehouse *w)
{
    for (int i = 0; i < buf->len-1; i++) // skip last line
    {
        read_line_to_warehouse(buf->buffer[i], w);
    }
}

void solve_1(const char *file_name, int piles_num)
{

    FILE * fp;

    fp = fopen(file_name, "r");
    if (fp == NULL) exit(EXIT_FAILURE);
    char chunk[CHUNK_SIZE];

    int counter = 0;
    bool read_piles = false;
    text_buffer text = new_text_buffer();
    int move, from, to;
    warehouse w = new_warehouse();

    while(fgets(chunk, sizeof(chunk), fp) != NULL)
    {
        if (strcmp(chunk, "\n") == 0)
        {
            read_piles = true;
            append_to_warehouse(&text, &w);
            continue;
        }
        if (!read_piles)
        {
            append_text_buffer(&text, chunk);
            continue;
        }
        sscanf(chunk, "move %i from %i to %i", &move, &from, &to);
        single_move_piles_warehouse(&w, move, from, to);
    }
    char *c = get_top_chars_in_warehouse(&w, piles_num);
    printf("pile arrangement to task 5.1 for data from %s is: %s \n", file_name, c);
    free(c);
    free_text_buffer(&text);
    free_warehouse(&w);
}

void solve_2(const char *file_name, int piles_num)
{

    FILE * fp;

    fp = fopen(file_name, "r");
    if (fp == NULL) exit(EXIT_FAILURE);
    char chunk[CHUNK_SIZE];

    int counter = 0;
    bool read_piles = false;
    text_buffer text = new_text_buffer();
    int move, from, to;
    warehouse w = new_warehouse();

    while(fgets(chunk, sizeof(chunk), fp) != NULL)
    {
        if (strcmp(chunk, "\n") == 0)
        {
            read_piles = true;
            append_to_warehouse(&text, &w);
            continue;
        }
        if (!read_piles)
        {
            append_text_buffer(&text, chunk);
            continue;
        }
        sscanf(chunk, "move %i from %i to %i", &move, &from, &to);
        multiple_move_piles_warehouse(&w, move, from, to);
    }

    char *c = get_top_chars_in_warehouse(&w, piles_num);
    printf("pile arrangement to task 5.2 for data from %s is: %s \n", file_name, c);
    free(c);
    free_text_buffer(&text);
    free_warehouse(&w);
}

int main()
{
    const char *TEST_FILE_NAME = "data_test.txt";
    solve_1(TEST_FILE_NAME, 3);
    solve_2(TEST_FILE_NAME, 3);

    const char *TASK_FILE_NAME = "data_task.txt";
    solve_1(TASK_FILE_NAME, 9);
    solve_2(TASK_FILE_NAME, 9);
}