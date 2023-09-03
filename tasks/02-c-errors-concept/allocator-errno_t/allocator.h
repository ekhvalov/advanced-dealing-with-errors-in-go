#include <errno.h>

#define __STDC_WANT_LIB_EXT1__ 1
// Если define выше не работает для нашего компилятора, то определяем тип руками:
// typedef int errno_t;

extern int errno;

#define ADMIN 777
#define MIN_MEMORY_BLOCK 1024

errno_t allocate(int user_id, size_t size, void **mem)
{
        errno_t err = 0;
        if (user_id != ADMIN) {
            err = EPERM;
        } else if (size < MIN_MEMORY_BLOCK) {
            err = EDOM;
        } else {
            *mem = malloc(size);
            if (*mem == NULL) {
                err = ENOMEM;
            }
        }
        if (err != 0) {
            mem = NULL;
        }
        return err;
}
