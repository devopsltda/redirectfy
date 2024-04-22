import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { trigger, state, style, transition, animate } from '@angular/animations';

@NgModule({
  declarations: [],
  imports: [
    CommonModule,
  ]
})
export class AnimationsModule { }

export const fadeInOutAnimation = [
  trigger('fade1000ms', [
  transition(':enter', [style({ opacity: 0 }), animate('1000ms', style({ opacity: 1 }))]),
  transition(':leave', [style({ opacity: 1 }), animate('1000ms', style({ opacity: 0 }))]),
])
,
 trigger('fade500ms', [
  transition(':enter', [style({ opacity: 0 }), animate('500ms', style({ opacity: 1 }))]),
  transition(':leave', [style({ opacity: 1 }), animate('500ms', style({ opacity: 0 }))]),
]),
trigger('fadeIn500ms', [
  transition(':enter', [style({ opacity: 0 }), animate('500ms', style({ opacity: 1 }))]),
])

];

export const SlideAnimation = [

  trigger('slideUpDown', [
  state('void', style({ height: '0', overflow: 'hidden' })), // Estado inicial quando oculto
  state('*', style({ height: '*', overflow: 'hidden' })), // Estado final quando visível
  transition('void => *', [
    animate('200ms ease-in', style({ height: '*', overflow: 'hidden' })),
  ]),
  transition('* => void', [
    animate('200ms ease-out', style({ height: '0', overflow: 'hidden' })),
  ])
  ]),

];

export const NavBarAnimation = [
trigger('expandCollapse', [
  state('collapsed', style({ width: '90px' })), // Estado inicial colapsado
  state('expanded', style({ width: '*' })), // Estado expandido
  transition('collapsed => expanded', animate('300ms ease-in')),
  transition('expanded => collapsed', animate('300ms ease-out')),
])





]
