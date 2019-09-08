module Lib.Sdl2 where

import Prelude (Unit)
import Effect (Effect)

foreign import doit :: Effect Unit

foreign import data InitFlag :: Type
foreign import initTimer :: InitFlag
foreign import initAudio :: InitFlag
foreign import initVideo :: InitFlag
foreign import initJoystick :: InitFlag
foreign import initHaptic :: InitFlag
foreign import initGameController :: InitFlag
foreign import initEvents :: InitFlag
foreign import initNoParachute :: InitFlag
foreign import initSensor :: InitFlag
foreign import initEverything :: InitFlag

