#ifndef __IS_POINTER_H__
#define __IS_POINTER_H__
template <typename T>
class IsPointer {
  public:
    static const bool value = false;
};

template <typename T>
class IsPointer<T*> {
  public:
    static const bool value = true;
};

#endif
