#ifndef GO_GDAL_PROJ_H_
#define GO_GDAL_PROJ_H_

#define ACCEPT_USE_OF_DEPRECATED_PROJ_API_H
#include <proj_api.h>

#if defined(WIN32) || defined(WINDOWS) || defined(_WIN32) || defined(_WINDOWS)
#define PROJCAPICALL __declspec(dllexport)
#else
#define PROJCAPICALL
#endif

PROJCAPICALL char *transform(projPJ srcdefn, projPJ dstdefn, long point_count,
                             double *x, double *y, double *z);
PROJCAPICALL char *fwd(projPJ src, double *lng, double *lat);
PROJCAPICALL char *inv(projPJ dst, double *lng, double *lat);
PROJCAPICALL char *get_err(void);

#endif