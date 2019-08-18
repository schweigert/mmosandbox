godepgraph -nostdlib -horizontal github.com/schweigert/mmosandbox/domain/ | ruby deps_recolor.rb | dot -Tpng -o domain/deps.png

godepgraph -nostdlib -horizontal github.com/schweigert/mmosandbox/apps/wweb | ruby deps_recolor.rb | dot -Tpng -o apps/wweb/deps.png
godepgraph -nostdlib -horizontal github.com/schweigert/mmosandbox/apps/wauth | ruby deps_recolor.rb | dot -Tpng -o apps/wauth/deps.png
godepgraph -nostdlib -horizontal github.com/schweigert/mmosandbox/apps/wgame | ruby deps_recolor.rb | dot -Tpng -o apps/wgame/deps.png

godepgraph -nostdlib -horizontal github.com/schweigert/mmosandbox/clients/wclient | ruby deps_recolor.rb | dot -Tpng -o clients/wclient/deps.png