- #+BEGIN_QUERY
  {
    :title [:b "Find blocks by Page.type=FrameModel"]
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
		- The same goes for `American manufacturer`, `American`, [[American manufacturer]], and [[American]]
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