IF(UNIX)
  SET(CMAKE_C_FLAGS "${CMAKE_C_FLAGS} -Wno-float-conversion -Wno-shadow -Wno-unused-function -Wno-format-extra-args -Wno-unknown-pragmas -Wno-tautological-compare -Wno-attributes -Wno-unused-const-variable -Wno-sign-compare -Wno-deprecated-declarations")
IF("${CMAKE_CXX_COMPILER_ID}" STREQUAL "GNU")
SET(CMAKE_C_FLAGS "${CMAKE_C_FLAGS} -Wno-maybe-uninitialized -Wno-misleading-indentation ")
ENDIF()
ENDIF()

if(UNIX AND NOT APPLE)
SET(CMAKE_C_FLAGS "${CMAKE_C_FLAGS} -fPIC")
ENDIF()


INCLUDE_DIRECTORIES(BEFORE 
  ${CMAKE_CURRENT_SOURCE_DIR}/proj/include
  ${CMAKE_CURRENT_SOURCE_DIR}/proj/src
)
 
INCLUDE_DIRECTORIES(
                    "${CMAKE_CURRENT_SOURCE_DIR}/../libsqlite3" 
                    "${CMAKE_CURRENT_BINARY_DIR}/../libsqlite3"
                    )

IF (WIN32)
  INCLUDE_DIRECTORIES(BEFORE ${CMAKE_CURRENT_SOURCE_DIR}/arch/win/src)
  set(CONFIG_SRC ${CMAKE_CURRENT_SOURCE_DIR}/arch/win/src/proj_config.h)
ELSE()
  INCLUDE_DIRECTORIES(BEFORE ${CMAKE_CURRENT_SOURCE_DIR}/arch/unix/src)
  set(CONFIG_SRC ${CMAKE_CURRENT_SOURCE_DIR}/arch/unix/src/proj_config.h)
ENDIF()
 
add_definitions( -DMUTEX_stub)
SET(SRC_LIBPROJ_PJ
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/4D_api.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/aasincos.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/adjlon.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/auth.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/ctx.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/datums.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/datum_set.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/deriv.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/dmstor.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/ellps.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/ell_set.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/errno.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/factors.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/fileapi.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/filemanager.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/fwd.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/gauss.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/generic_inverse.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/grids.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/initcache.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/init.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/internal.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/inv.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/list.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/log.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/malloc.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/mlfn.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/msfn.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/mutex.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/networkfilemanager.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/param.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/phi2.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/pipeline.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/pr_list.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/proj_json_streaming_writer.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/proj_mdist.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/qsfn.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/release.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/rtodms.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/sqlite3_utils.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/strerrno.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/strtod.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/tracing.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/transform.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/tsfn.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/units.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/utils.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/wkt1_parser.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/wkt2_parser.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/wkt_parser.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/zpoly1.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/geodesic.c
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/wkt1_generated_parser.c
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/wkt2_generated_parser.c
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/apps/cct.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/apps/cs2cs.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/apps/emess.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/apps/geod.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/apps/geod_interface.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/apps/geod_set.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/apps/gie.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/apps/proj.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/apps/projinfo.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/apps/proj_strtod.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/apps/projsync.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/apps/utils_apps.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/conversions/axisswap.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/conversions/cart.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/conversions/geoc.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/conversions/geocent_conversions.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/conversions/noop.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/conversions/set.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/conversions/unitconvert.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/iso19111/c_api.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/iso19111/common.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/iso19111/coordinateoperation.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/iso19111/coordinatesystem.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/iso19111/crs.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/iso19111/datum.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/iso19111/factory.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/iso19111/internal_iso19111.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/iso19111/io.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/iso19111/metadata.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/iso19111/static.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/iso19111/util.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/adams.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/aea.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/aeqd.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/airy.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/aitoff.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/august.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/bacon.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/bertin1953.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/bipc.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/boggs.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/bonne.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/calcofi.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/cass.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/cc.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/ccon.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/cea.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/chamb.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/collg.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/col_urban.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/comill.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/crast.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/denoy.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/eck1.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/eck2.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/eck3.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/eck4.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/eck5.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/eqc.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/eqdc.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/eqearth.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/fahey.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/fouc_s.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/gall.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/geos.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/gins8.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/gnom.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/gn_sinu.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/goode.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/gstmerc.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/hammer.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/hatano.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/healpix.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/igh.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/igh_o.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/imw_p.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/isea.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/krovak.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/labrd.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/laea.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/lagrng.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/larr.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/lask.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/latlong.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/lcca.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/lcc.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/loxim.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/lsat.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/mbtfpp.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/mbtfpq.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/mbt_fps.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/merc.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/mill.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/misrsom.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/mod_ster.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/moll.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/natearth2.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/natearth.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/nell.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/nell_h.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/nicol.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/nsper.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/nzmg.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/ob_tran.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/ocea.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/oea.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/omerc.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/ortho.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/patterson.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/poly.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/putp2.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/putp3.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/putp4p.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/putp5.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/putp6.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/qsc.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/robin.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/rouss.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/rpoly.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/sch.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/sconics.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/somerc.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/sterea.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/stere.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/sts.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/tcc.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/tcea.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/times.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/tmerc.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/tobmerc.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/tpeqd.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/urm5.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/urmfps.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/vandg2.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/vandg4.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/vandg.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/wag2.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/wag3.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/wag7.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/wink1.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/projections/wink2.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/transformations/affine.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/transformations/defmodel.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/transformations/deformation.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/transformations/helmert.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/transformations/hgridshift.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/transformations/horner.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/transformations/molodensky.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/transformations/tinshift.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/transformations/vgridshift.cpp
${CMAKE_CURRENT_SOURCE_DIR}/proj/src/transformations/xyzgridshift.cpp
)

set(HEADERS_LIBPROJ
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/include/proj/common.hpp
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/include/proj/coordinateoperation.hpp
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/include/proj/coordinatesystem.hpp
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/include/proj/crs.hpp
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/include/proj/datum.hpp
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/include/proj/io.hpp
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/include/proj/metadata.hpp
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/include/proj/nn.hpp
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/include/proj/util.hpp

        ${CMAKE_CURRENT_SOURCE_DIR}/proj/include/proj/internal/coordinateoperation_constants.hpp
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/include/proj/internal/coordinateoperation_internal.hpp
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/include/proj/internal/coordinatesystem_internal.hpp
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/include/proj/internal/esri_projection_mappings.hpp
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/include/proj/internal/include_nlohmann_json.hpp
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/include/proj/internal/internal.hpp
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/include/proj/internal/io_internal.hpp
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/include/proj/internal/lru_cache.hpp
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/include/proj/internal/tracing.hpp
        ${CMAKE_CURRENT_SOURCE_DIR}/proj/include/proj/internal/nlohmann/json.hpp
)

source_group("Header Files" FILES ${HEADERS_LIBPROJ})
source_group("Source Files" FILES ${SRC_LIBPROJ_PJ})
include_directories( ${CMAKE_CURRENT_BINARY_DIR})

add_definitions(-DPROJ_LIB="${CMAKE_INSTALL_PREFIX}/${DATADIR}")

set(ALL_LIBPROJ_SOURCES ${SRC_LIBPROJ_PJ})
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
