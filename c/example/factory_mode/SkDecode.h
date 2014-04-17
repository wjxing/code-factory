#include <stdlib.h>
#include "SkStream.h"
namespace skia {
class SkDecode {
  public:
    virtual void decode() = 0;
    static SkDecode* factory(SkStream* stream);
};
};
