#ifndef __CONTAINER_OF_H__
#define __CONTAINER_OF_H__

#if 0
#define offset(type, member) ((size_t)(&((type *)0)->member))

#define container_of(ptr, type, member) ({ \
        const typeof(((type *)0)->member) *__mptr = (ptr); \
        (type *)(((char *)__mptr) - offset(type, member)); \
    })
#else
#define container_of(ptr, type, member) ((type *)((char *)ptr - (size_t)(&((type *)0)->member)))
#endif

#endif
