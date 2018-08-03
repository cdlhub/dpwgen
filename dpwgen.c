// #define _GNU_SOURCE
#include <errno.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
// #include <gsl/gsl_rng.h>

const int MAX_NUM_DICE = 5;
const int SIX_POW[MAX_NUM_DICE] = { 1, 6, 36, 216, 1296 };   

#define FREE_TO_NULL(p) \
{ \
    free(p); \
    p = NULL; \
}

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

int compute_line(const int d[], int num_dice)
{
    int line = 1;
    for (int i = 0; i < num_dice; i++)
    {
        line += (d[i]-1) * SIX_POW[i];
    }
    return line;
}

int read_line(char **buf, const char *file_name, int line_number)
{
    FILE *file = fopen(file_name, "r");
    if (file == NULL)
    {
        return -1;
    }

    size_t buf_size = 0;
    int len = 0;
    for (int count = 1; (len = getline(buf, &buf_size, file)) != -1; count++)
    {
        if (count == line_number)
        {
            break;
        }

        // init for next getline()
        FREE_TO_NULL(*buf);
        buf_size = 0;
    }
    fclose(file);

    // getline() error or file contains less lines than line_number
    if (len == -1)
    {
        FREE_TO_NULL(*buf);
        return -1;
    }

    return len;
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

int main(int argc, char const *argv[])
{
    int d[MAX_NUM_DICE] = { 0, 0, 0, 0, 0 };

    for (int i = 0; i < MAX_NUM_DICE; i++)
    {
        d[i] = scan_dice_value(i+1);
    }

    int line_number = compute_line(d, MAX_NUM_DICE);
    printf("line:     %d\n", line_number);

    char *file_name = "eff_large_wordlist.txt";
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
