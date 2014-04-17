namespace skia {
class SkStream {
  public:
    SkStream(const char* stream);
    virtual ~SkStream();
    const char* read();
  private:
    const char* mpStream;
};
};
