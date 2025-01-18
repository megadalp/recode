# recode для Linux, проверено на KDE Plasma X11
# Исправление текста, набранного когда забыл переключить раскладку :)

## Оригинал на perl:
    
    use utf8;
    $lat=q(`~!@#$%^&qwertyuiop[]asdfghjkl;'zxcvbnm,./QWERTYUIOP{}ASDFGHJKL:"|ZXCVBNM<>?);
    $cyr=q(ёЁ!"№;%:?йцукенгшщзхъфывапролджэячсмитьбю.ЙЦУКЕНГШЩЗХЪФЫВАПРОЛДЖЭ/ЯЧСМИТЬБЮ,);
    while(<>) { eval "tr{$lat$cyr}{$cyr$lat}"; print; }

## Настроить исправление раскладки набранного (выделенного) текста.
- проверить наличие/установить утилиты
    xdotool
    xclip
- Создать бинарник recode
    - на perl (код выше) или на go (бинарник recode готов к использованию)
    - поместить его например в /usr/local/bin
- Зайти: Настройки / Комбинация клавиш / Специальные действия
    - добавить группу [dalp], внутри нее создать действие
       - Комментарий типа
           recode to another charset
       - Действие
           xdotool key 'ctrl+c'; xclip -out -sel clip | recode | xclip -in -sel clip; xdotool key 'ctrl+v'
