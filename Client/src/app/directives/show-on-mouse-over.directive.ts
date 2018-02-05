import { Directive } from '@angular/core';
import {
  ElementRef,
  HostListener,
  Renderer2,
  Input,
  TemplateRef,
  ViewContainerRef
} from '@angular/core';

@Directive({
  selector: '[appShowOnMouseOver]'
})
export class ShowOnMouseOverDirective {

  @Input() template: TemplateRef<any>;

  constructor(private el: ElementRef,
    private renderer: Renderer2,
    private viewContainer: ViewContainerRef) {
    console.log(this.el.nativeElement)
  }

  @HostListener('mouseover') onMouseOver() {
    console.log('MouseOver called...');
    console.log(this.template);
    //this.renderer.setStyle(this.el.nativeElement, 'display', 'block');
    this.renderer.setStyle(this.template, 'display', 'block');
  }

  @HostListener('mouseleave') onMouseLeave() {
    console.log('MouseLeave called...');
    //this.renderer.setStyle(this.el.nativeElement, 'display', 'none');
    // this.viewContainer.clear();
    this.renderer.setStyle(this.template, 'display', 'none');
  }

}
