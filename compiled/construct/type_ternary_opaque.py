from construct import *
from construct.lib import *

type_ternary_opaque = Struct(
	'dif_wo_hack' / ???,
	'dif_with_hack' / ???,
)

_schema = type_ternary_opaque
