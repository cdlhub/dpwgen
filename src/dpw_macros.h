#ifndef _DPW_MACROS_H_
#define _DPW_MACROS_H_

#include <stdlib.h>

#define FREE_TO_NULL(p) \
{ \
    free(p); \
    (p) = NULL; \
}

#endif // _DPW_MACROS_H_
