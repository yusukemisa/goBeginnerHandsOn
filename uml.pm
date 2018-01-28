@startuml
title おみくじフローチャート

start

:1~6までランダムに数字作る;
:生成した数字を判定;
if (1) then
:凶を表示;
else if (2,3) then
:吉を表示;
else if (4,5) then
:中吉を表示;
else if (6) then
:大吉を表示;
else()
:エラー;
stop
endif

stop

@enduml