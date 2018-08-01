#include <gsl/gsl_rng.h>

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

int main(int argc, char const *argv[])
{
    int d[MAX_NUM_DICE] = { 0, 0, 0, 0, 0 };

    for (int i = 0; i < MAX_NUM_DICE; i++)
    {
        d[i] = scan_dice_value(i+1);
    }

    int line = compute_line(d, MAX_NUM_DICE);
    printf("password line is %d\n", line);

    return 0;
}
