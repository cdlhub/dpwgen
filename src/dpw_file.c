#include <stdio.h>
#include <string.h>

#include "dpw_macros.h"

const int SIX_POW_SIZE = 5;
const int SIX_POW[] = { 1, 6, 36, 216, 1296 };

int compute_line(const int d[], int num_dice)
{
    // assert(num_dice < SIX_POW_SIZE)

    int line = 1;
    for (int i = 0; i < num_dice; i++)
    {
        line += (d[i]-1) * SIX_POW[i];
    }
    return line;
}

ssize_t read_line(char **buf, const char *file_name, int line_number)
{
    FILE *file = fopen(file_name, "r");
    if (file == NULL)
    {
        return -1;
    }

    size_t buf_size = 0;
    ssize_t len = 0;
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
    }

    return len;
}

char* get_password(const char *buf, size_t len, int num_dice)
{
    int start = num_dice + 1;
    size_t size = len - start;
    if (buf[len-1] == '\n')
    {
        size--;
    }

    char* pwd = malloc(size + 1);
    memcpy(pwd, buf + start, size);
    pwd[size] = '\0';
    return pwd;
}
