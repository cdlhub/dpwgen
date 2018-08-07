#define _GNU_SOURCE
#include <errno.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
// #include <gsl/gsl_rng.h>

#include "dpw_file.h"
#include "dpw_macros.h"

const int MAX_NUM_DICE = 5;

int scan_dice_value(int n)
{
    int val = 0;

    printf("Enter value for D%d in range [1, 6]: ", n);
    scanf("%d", &val);
    while (val < 1 || 6 < val)
    {
        while (getchar() != '\n');
        fprintf(stderr, "error: input value bad format: value must be an integer in [1, 6] range\n");
        printf("Enter value for D%d: ", n);
        scanf("%d", &val);
    }

    return val;
}

void print_password(const char *buf, int len)
{
    int start = MAX_NUM_DICE + 1;
    int size = len - start;
    if (buf[len-1] == '\n')
    {
        size--;
    }

    printf("password: '%.*s'\n", size, buf + start);
}

int main(void)
{
    int d[] = { 0, 0, 0, 0, 0 };


    for (int i = 0; i < MAX_NUM_DICE; i++)
    {
        d[i] = scan_dice_value(i+1);
    }

    int line_number = compute_line(d, MAX_NUM_DICE);
    printf("line:     %d\n", line_number);

    char *file_name = "../eff_large_wordlist.txt";
    // char *file_name = "../empty.txt";
    char *password = NULL;
    int password_len = read_line(&password, file_name, line_number);
    if (password_len == -1)
    {
        fprintf(stderr, "error: cannot read password from '%s'", file_name);
        return 1;
    }

    print_password(password, password_len);
    free(password);

    return 0;
}
