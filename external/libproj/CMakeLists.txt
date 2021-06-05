IF(UNIX)
  SET(CMAKE_C_FLAGS "${CMAKE_C_FLAGS} -Wno-float-conversion -Wno-shadow -Wno-unused-function -Wno-format-extra-args -Wno-unknown-pragmas -Wno-tautological-compare -Wno-attributes -Wno-unused-const-variable -Wno-sign-compare -Wno-deprecated-declarations")
IF("${CMAKE_CXX_COMPILER_ID}" STREQUAL "GNU")
SET(CMAKE_C_FLAGS "${CMAKE_C_FLAGS} -Wno-maybe-uninitialized -Wno-misleading-indentation ")
ENDIF()
ENDIF()

if(UNIX AND NOT APPLE)
SET(CMAKE_C_FLAGS "${CMAKE_C_FLAGS} -fPIC")
ENDIF()


INCLUDE_DIRECTORIES(BEFORE ${CMAKE_CURRENT_SOURCE_DIR}/proj/src)

IF (WIN32)
  INCLUDE_DIRECTORIES(BEFORE ${CMAKE_CURRENT_SOURCE_DIR}/arch/win/src)
  set(CONFIG_SRC ${CMAKE_CURRENT_SOURCE_DIR}/arch/win/src/proj_config.h)
ELSE()
  INCLUDE_DIRECTORIES(BEFORE ${CMAKE_CURRENT_SOURCE_DIR}/arch/unix/src)
  set(CONFIG_SRC ${CMAKE_CURRENT_SOURCE_DIR}/arch/unix/src/proj_config.h)
ENDIF()
 
add_definitions( -DMUTEX_stub)
SET(SRC_LIBPROJ_PJ
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/nad_init.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_aea.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_aeqd.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_airy.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_aitoff.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_august.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_bacon.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_bipc.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_boggs.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_bonne.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_calcofi.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_cass.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_cc.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_cea.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_chamb.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_collg.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_crast.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_denoy.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_eck1.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_eck2.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_eck3.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_eck4.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_eck5.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_eqc.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_eqdc.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_fahey.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_fouc_s.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_gall.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_geos.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_gins8.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_gnom.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_gn_sinu.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_goode.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_gstmerc.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_hammer.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_hatano.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_igh.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_isea.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_imw_p.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_krovak.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_labrd.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_laea.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_lagrng.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_larr.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_lask.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_lcca.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_lcc.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_loxim.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_lsat.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_mbt_fps.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_mbtfpp.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_mbtfpq.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_merc.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_mill.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_mod_ster.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_moll.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_natearth.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_nell.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_nell_h.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_nocol.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_nsper.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_nzmg.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_ob_tran.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_ocea.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_oea.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_omerc.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_ortho.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_poly.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_putp2.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_putp3.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_putp4p.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_putp5.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_putp6.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_qsc.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_robin.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_rpoly.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_sconics.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_somerc.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_sterea.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_stere.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_sts.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_tcc.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_tcea.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_tmerc.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_tpeqd.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_urm5.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_urmfps.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_vandg.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_vandg2.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_vandg4.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_wag2.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_wag3.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_wag7.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_wink1.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_wink2.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/proj_etmerc.c
)

SET(SRC_LIBPROJ_CORE
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/aasincos.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/adjlon.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/bch2bps.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/bchgen.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/biveval.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/dmstor.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/emess.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/emess.h
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/geocent.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/geocent.h
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/geodesic.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/mk_cheby.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/nad_cvt.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/nad_init.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/nad_intr.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/pj_apply_gridshift.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/pj_apply_vgridshift.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/pj_auth.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/pj_ctx.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/pj_fileapi.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/pj_datum_set.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/pj_datums.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/pj_deriv.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/pj_ell_set.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/pj_ellps.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/pj_errno.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/pj_factors.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/pj_fwd.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/pj_gauss.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/pj_gc_reader.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/pj_geocent.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/pj_gridcatalog.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/pj_gridinfo.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/pj_gridlist.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/PJ_healpix.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/pj_init.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/pj_initcache.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/pj_inv.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/pj_latlong.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/pj_list.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/pj_list.h
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/pj_log.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/pj_malloc.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/pj_mlfn.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/pj_msfn.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/pj_mutex.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/pj_open_lib.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/pj_param.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/pj_phi2.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/pj_pr_list.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/pj_qsfn.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/pj_release.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/pj_strerrno.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/pj_transform.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/pj_tsfn.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/pj_units.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/pj_utils.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/pj_zpoly1.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/proj_mdist.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/proj_rouss.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/rtodms.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/vector1.c
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/pj_strtod.c
        ${CONFIG_SRC}
 )

set(HEADERS_LIBPROJ
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projects.h
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/proj_api.h
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/src/geodesic.h
)

source_group("Header Files" FILES ${HEADERS_LIBPROJ})
source_group("Source Files\\Core" FILES ${SRC_LIBPROJ_CORE})
source_group("Source Files\\PJ" FILES ${SRC_LIBPROJ_PJ})
include_directories( ${CMAKE_CURRENT_BINARY_DIR})

add_definitions(-DPROJ_LIB="${CMAKE_INSTALL_PREFIX}/${DATADIR}")

set(ALL_LIBPROJ_SOURCES ${SRC_LIBPROJ_PJ} ${SRC_LIBPROJ_CORE})
set(ALL_LIBPROJ_HEADERS ${HEADERS_LIBPROJ} )

ADD_LIBRARY(proj STATIC
  ${ALL_LIBPROJ_SOURCES}
  ${ALL_LIBPROJ_HEADERS}
          )

TARGET_INCLUDE_DIRECTORIES(proj PUBLIC $<BUILD_INTERFACE:${CMAKE_CURRENT_SOURCE_DIR}/..>)

SOURCE_GROUP("src" FILES ${ALL_LIBPROJ_SOURCES})
SOURCE_GROUP("include" FILES ${ALL_LIBPROJ_HEADERS})

SET_TARGET_PROPERTIES(proj
                    PROPERTIES PUBLIC_HEADER "${headers}")
SET_TARGET_PROPERTIES(proj PROPERTIES
  ARCHIVE_OUTPUT_DIRECTORY_DEBUG ${CMAKE_CURRENT_BINARY_DIR}
  ARCHIVE_OUTPUT_DIRECTORY_RELEASE ${CMAKE_CURRENT_BINARY_DIR})
SET_TARGET_PROPERTIES(proj PROPERTIES 
  LIBRARY_OUTPUT_DIRECTORY_DEBUG ${CMAKE_CURRENT_BINARY_DIR}
  LIBRARY_OUTPUT_DIRECTORY_RELEASE ${CMAKE_CURRENT_BINARY_DIR})
SET_TARGET_PROPERTIES(proj PROPERTIES
  RUNTIME_OUTPUT_DIRECTORY_DEBUG ${CMAKE_CURRENT_BINARY_DIR}
  RUNTIME_OUTPUT_DIRECTORY_RELEASE ${CMAKE_CURRENT_BINARY_DIR})

IF(FLYWAVE_ENABLE_SOLUTION_FOLDERS)
  SET_TARGET_PROPERTIES(proj PROPERTIES FOLDER external)
ENDIF(FLYWAVE_ENABLE_SOLUTION_FOLDERS)
