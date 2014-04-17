#ifndef __SHOWMACRO_H__
#define __SHOWMACRO_H__
#define MACRO_(x)       #x
// This macro could show THE mac content
#define SHOWMACRO(mac)  MACRO_(mac)
#endif
