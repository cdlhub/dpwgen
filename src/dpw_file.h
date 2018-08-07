#ifndef _DPW_FILE_H_
#define _DPW_FILE_H_

/*! 
 * @brief Compute line number of the password based on dice rolls
 * @pre d is n array of num_dice values in range [1, 6].
 * @desc If d = { 1, 5, 3, 3, 2 }, then line number is:
 *       @f$ \sum_{i=0}^{num_dice-1} (d[i]-1) * 6^i @f$
 * @return The line number in password file where for values in d.
 */
int compute_line(const int d[], int num_dice);

/*!
 * buf is allocated in the function. You are responsible to call free after use.
 */
int read_line(char **buf, const char *file_name, int line_number);

#endif // _DPW_FILE_H_
