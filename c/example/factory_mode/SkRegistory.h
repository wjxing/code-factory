namespace skia {
template <typename Decode, typename Stream>
class SkRegistory {
  public:
    typedef Decode (*Factory)(Stream);
    SkRegistory(Factory fact)
        : fFact(fact) {
        SkRegistory* reg = gHead;
        while (reg) {
            if (reg == this)
                return;
            reg = reg->fChain;
        }
        fChain = gHead;
        gHead = this;
    }

    static const SkRegistory* head() {
        return gHead;
    }

    const SkRegistory* next() const {
        return fChain;
    }

    Factory factory() const {
        return fFact;
    }

  private:
    Factory fFact;
    SkRegistory* fChain;
    static SkRegistory* gHead;
};

template <typename Decode, typename Stream>
SkRegistory<Decode, Stream>* SkRegistory<Decode, Stream>::gHead;
};
