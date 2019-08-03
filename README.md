# mmosandbox
MMORPG Playground

## Dev

## Deps

Setup:

```
    go get github.com/kisielk/godepgraph
    sudo pacman -S graphviz
```

```
    make deps
```

## TODO

 - [x] Criar regra de negócio para criar conta
 - [x] Criar cliente para criar conta na arq Willson
 - [ ] Criar cliente para criar conta na arq Salz
 - [ ] Criar cliente para criar conta na arq Rudy
 - [ ] Criar Regra de negócio para criar personagem
 - [ ] Criar cliente para criar personagem na arq Willson
 - [ ] Criar cliente para criar personagem na arq Salz
 - [ ] Criar cliente para criar personagem na arq Rudy

## Domain

![deps.png](domain/deps.png)

## Services

### apps/wauth

![deps.png](apps/wauth/deps.png)

### apps/wgame

![deps.png](apps/wgame/deps.png)

### apps/wweb

![deps.png](apps/wweb/deps.png)

## Clients

### clients/wclient

![deps.png](clients/wclient/deps.png)