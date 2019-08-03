godepgraph -nostdlib -horizontal github.com/schweigert/mmosandbox/domain/ | dot -Tpng -o domain/deps.png

godepgraph -nostdlib -horizontal github.com/schweigert/mmosandbox/apps/wweb | dot -Tpng -o apps/wweb/deps.png
godepgraph -nostdlib -horizontal github.com/schweigert/mmosandbox/apps/wauth | dot -Tpng -o apps/wauth/deps.png
godepgraph -nostdlib -horizontal github.com/schweigert/mmosandbox/apps/wgame | dot -Tpng -o apps/wgame/deps.png

godepgraph -nostdlib -horizontal github.com/schweigert/mmosandbox/clients/wclient | dot -Tpng -o clients/wclient/deps.png