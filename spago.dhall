{-
Welcome to a Spago project!
You can edit this file as you like.
-}
{ name =
    "PureScript Game Thing"
, dependencies =
    [ "effect", "console", "psci-support" ]
, packages =
    ./packages.dhall
, sources =
    [ "src/**/*.purs", "test/**/*.purs" ]
}
