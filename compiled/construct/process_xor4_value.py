from construct import *
from construct.lib import *

process_xor4_value = Struct(
	'key' / ???,
	'buf' / ???,
)

_schema = process_xor4_value
