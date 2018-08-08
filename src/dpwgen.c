#define _GNU_SOURCE
#include <stdio.h>
#include <stdlib.h>
#include <time.h>
// #include <gsl/gsl_rng.h>

#include "dpw_file.h"
#include "dpw_macros.h"

#define VERSION "0.2.0"

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

void init_draw(unsigned int seed)
{
    srand(seed);
}

void draw(int d[], size_t num_dice)
{
    for (size_t i = 0; i < num_dice; i++)
    {
        d[i] = (rand() % 6) + 1;
    }
}

int main(void)
{
    const size_t num_dice = 5;
    const int num_words = 128;
    const char* const pwd_file_name = "../eff_large_wordlist.txt";

    // print header
    printf("| line | dice draw | password     |\n");
    printf("| ---- | --------- | ------------ |\n");

    init_draw((unsigned) time(NULL));
    int d[num_dice];
    for (int i = 0; i < num_words; i++)
    {
        draw(d, num_dice);

        char* pwd_line;
        int pwd_line_num = compute_line(d, num_dice);
        ssize_t pwd_len = read_line(&pwd_line, pwd_file_name, pwd_line_num);
        if (pwd_len == -1)
        {
            fprintf(stderr, "error: cannot read password from '%s'", pwd_file_name);
            return EXIT_FAILURE;
        }

        char* pwd = get_password(pwd_line, pwd_len, num_dice);
        FREE_TO_NULL(pwd_line);

        printf("| %4d |     %d%d%d%d%d | %-12s |\n", pwd_line_num, d[0], d[1], d[2], d[3], d[4], pwd);
        FREE_TO_NULL(pwd);
    }

    return EXIT_SUCCESS;
}
