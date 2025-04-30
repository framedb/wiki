- #+BEGIN_QUERY
  {
    :title [:b "WIP Find blocks by Page.type=FrameModel"]
    :inputs [:frame-bridge "Keyhole bridge"]
    :query [
      :find (pull ?b [*])
        :in $ ?key ?val
        :where
        [?b :block/page ?p]
        [?b :block/parent ?parent]
        (page-property ?p :frame-origin "American")
        (property ?parent :type "features")
        (property ?b ?key ?val)
    ]
   }
  #+END_QUERY
	- TODO: Follow reference when checking for enum in props
		- `"Keyhole bridge"` `"keyhole"`, [[keyhole]], and [[Keyhole bridge]] should work all the same
		- The same goes for `American manufacturer`, `American`, [[American]], and [[American]]
-
- #+BEGIN_QUERY
  {
    :title [:b "WIP Find blocks with alias"]
    :inputs [:frame-bridge "Keyhole page"]
    :query [
      :find (pull ?b [*])
        :in $ ?key ?val
        :where
        [?b :block/page ?p]
        [?b :block/parent ?parent]
        (page-property ?p :frame-origin "American")
        (property ?parent :type "features")
        (?b :block/properties ?props)
        [(get ?props :frame-bridge) ?bridge]
        (or-join [?b ?val ?bridge]
           [?b :block/refs ?val]
           (and
              [?p :block/alias ?a]
              [?b :block/refs ?a]
           )
        )
    ]
   }
  #+END_QUERY
-
- #+BEGIN_QUERY
  {
    :title [:b "Find features for all frame pages"]
    :inputs [1 2]
    :query [
      :find (pull ?b [*])
        :in $ ?key ?val
        :where
        [?b :block/page ?p]
        [?b :block/parent ?parent]
        (page-property ?p :type "FrameModel" :frame-origin "America")
        (property ?b :type "features")
    ]
   }
  #+END_QUERY