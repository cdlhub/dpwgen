// #define _GNU_SOURCE
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
// #include <gsl/gsl_rng.h>

const int MAX_NUM_DICE = 5;
const int SIX_POW[MAX_NUM_DICE] = { 1, 6, 36, 216, 1296 };   

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

char* read_line(const char *file_name, int line_number)
{
    FILE *f = fopen(file_name, "r");
    if (f == NULL)
    {
        return NULL;
    }

    size_t n = 0;
    char *buf = NULL;
    int len = 0;
    // TODO:
    //  - manage getline error (=> free(buf))
    //  - manage empty lines
    for (int count = 1; (len = getline(&buf, &n, f)) >= 0; count++)
    {
        if (count == line_number)
        {
            break;
        }

        free(buf);
        buf = NULL;
        n = 0;
    }
    fclose(f);

    return buf;
}

int main(int argc, char const *argv[])
{
    int d[MAX_NUM_DICE] = { 0, 0, 0, 0, 0 };

    for (int i = 0; i < MAX_NUM_DICE; i++)
    {
        d[i] = scan_dice_value(i+1);
    }

    int line = compute_line(d, MAX_NUM_DICE);
    printf("line:     %d\n", line);

    char *file_name = "eff_large_wordlist.txt";
    char *password = read_line(file_name, line);
    if (password == NULL)
    {
        fprintf(stderr, "error: cannot read '%s'", file_name);
        return 1;
    }

    printf("password: '%s'\n", password+MAX_NUM_DICE+1);
    printf("length: %lu\n", strlen(password));
    free(password);

    return 0;
}
