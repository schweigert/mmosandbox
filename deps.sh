godepgraph -nostdlib -horizontal github.com/schweigert/mmosandbox/domain/ | ruby deps_recolor.rb | dot -Tpng -o domain/deps.png

godepgraph -nostdlib -horizontal github.com/schweigert/mmosandbox/apps/wweb | ruby deps_recolor.rb | dot -Tpng -o apps/wweb/deps.png
godepgraph -nostdlib -horizontal github.com/schweigert/mmosandbox/apps/wauth | ruby deps_recolor.rb | dot -Tpng -o apps/wauth/deps.png
godepgraph -nostdlib -horizontal github.com/schweigert/mmosandbox/apps/wgame | ruby deps_recolor.rb | dot -Tpng -o apps/wgame/deps.png

godepgraph -nostdlib -horizontal github.com/schweigert/mmosandbox/apps/sweb | ruby deps_recolor.rb | dot -Tpng -o apps/sweb/deps.png
godepgraph -nostdlib -horizontal github.com/schweigert/mmosandbox/apps/sauth | ruby deps_recolor.rb | dot -Tpng -o apps/sauth/deps.png
godepgraph -nostdlib -horizontal github.com/schweigert/mmosandbox/apps/sgame | ruby deps_recolor.rb | dot -Tpng -o apps/sgame/deps.png
godepgraph -nostdlib -horizontal github.com/schweigert/mmosandbox/apps/schat | ruby deps_recolor.rb | dot -Tpng -o apps/schat/deps.png

godepgraph -nostdlib -horizontal github.com/schweigert/mmosandbox/apps/rweb | ruby deps_recolor.rb | dot -Tpng -o apps/rweb/deps.png
godepgraph -nostdlib -horizontal github.com/schweigert/mmosandbox/apps/rcrud | ruby deps_recolor.rb | dot -Tpng -o apps/rcrud/deps.png
godepgraph -nostdlib -horizontal github.com/schweigert/mmosandbox/apps/rauth | ruby deps_recolor.rb | dot -Tpng -o apps/rauth/deps.png
godepgraph -nostdlib -horizontal github.com/schweigert/mmosandbox/apps/rgame | ruby deps_recolor.rb | dot -Tpng -o apps/rgame/deps.png

godepgraph -nostdlib -horizontal github.com/schweigert/mmosandbox/clients/wclient | ruby deps_recolor.rb | dot -Tpng -o clients/wclient/deps.png