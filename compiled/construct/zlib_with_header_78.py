from construct import *
from construct.lib import *

zlib_with_header_78 = Struct(
	'data' / GreedyBytes,
)

_schema = zlib_with_header_78
