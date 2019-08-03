godepgraph -withgoroot -d -nostdlib -horizontal github.com/schweigert/mmosandbox/apps/wweb | dot -y -Tpng -o apps/wweb/deps.png
godepgraph -withgoroot -d -nostdlib -horizontal github.com/schweigert/mmosandbox/apps/wauth | dot -y -Tpng -o apps/wauth/deps.png
godepgraph -withgoroot -d -nostdlib -horizontal github.com/schweigert/mmosandbox/apps/wgame | dot -y -Tpng -o apps/wgame/deps.png

godepgraph -withgoroot -d -nostdlib -horizontal github.com/schweigert/mmosandbox/clients/wclient | dot -y -Tpng -o clients/wclient/deps.png