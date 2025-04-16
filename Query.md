- #+BEGIN_QUERY
  {
    :title [:b "Find by 'parent has prop type:features' whose content contains some substring"]
    :query [
      :find (pull ?b [*])
        :where
        [?b :block/page ?p]
        [?b :block/parent ?parent]
        (property ?parent :type "features")
        [?b :block/content ?bcontent]
        [(clojure.string/includes? ?bcontent "frame-a")]
    ]
   }
  #+END_QUERY
-
- #+BEGIN_QUERY
  {
    :title [:b "Find block by Page.type=FrameModel with Feature.type=features"]
    :query [
      :find (pull ?b [*])
        :where
        [?b :block/page ?p]
        [?b :block/parent ?parent]
        (page-property ?p :type "FrameModel")
        (property ?parent :type "features")
        (property ?b :frame-rivet "Dot rivets")
    ]
   }
  #+END_QUERY
-
- #+BEGIN_QUERY
  {
    :title [:b "Find page by Page.type=FrameModel with Feature.type=features frame-rivet=[[Dot rivets]]"]
    :inputs [:frame-rivet "Dot rivets"]
    :query [
      :find (pull ?p [*])
        :in $ ?key ?val
        :where
        [?b :block/page ?p]
        [?b :block/parent ?parent]
        (page-property ?p :type "FrameModel")
        (property ?parent :type "features")
        (property ?b ?key ?val)
    ]
   }
  #+END_QUERY
-
- #+BEGIN_QUERY
  {
    :title [:b "Find by 'content contains prop name'"]
    :query [
      :find (pull ?b [*])
        :where
        [?b :block/page ?p]
        [?b :block/parent ?parent]
        (property ?parent :type "features")
        [?b :block/content ?bcontent]
        [(clojure.string/includes? ?bcontent "frame-rivet")]
    ]
   }
  #+END_QUERY
