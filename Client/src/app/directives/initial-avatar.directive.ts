import {
  Directive,
  ElementRef,
  Renderer2,
  Input,
  AfterViewInit,
  HostBinding
} from '@angular/core';


@Directive({
  selector: 'img[initialAvatar]'
})
export class InitialAvatarDirective implements AfterViewInit {

  private _pictureFormat = "png";
  private _fontScale = 100;
  private _fontWeight = 100;
  private _colorPalette = ["#bdc3c7",
    "#6f7b87",
    "#2c3e50",
    "#2f3193",
    "#662d91",
    "#922790",
    "#ec2176",
    "#ed1c24",
    "#f36622",
    "#f8941e",
    "#fab70f",
    "#fdde00",
    "#d1d219",
    "#8ec73f",
    "#00a650",
    "#00aa9c",
    "#00adef",
    "#0081cd",
    "#005bab"];

  constructor(private el: ElementRef,
    private renderer: Renderer2) {
    // console.log('Inital Avatar constructor call!!!!')
  }

  @Input() initial: string;
  @Input() bgColor: string = "grey";
  @Input() showShadow: boolean = true;
  @Input() textColor: string = "white";
  @Input() fontSize: number = 24;
  @Input() width: number = 40;
  @Input() height: number = 40;

  @HostBinding('src') private imgSrc: string;

  ngAfterViewInit(): void {
    //let avatarUrl = this.generateAvatar()
    //this.imgSrc = avatarUrl;
    //this.renderer.setAttribute(this.el.nativeElement, "src", avatarUrl);
    // this.renderer.setStyle(this.el.nativeElement, "background", "yellow");
  }

  ngOnChanges() {
    this.imgSrc = this.generateAvatar();
    this.renderer.addClass(this.el.nativeElement, "pixelated-image");
    // this.renderer.setStyle(this.el.nativeElement, "border-radius", `${this.width}px`);
  }

  private generateAvatar(): string {
    console.log('generating avatar...');
    // const width = 48, height = 48;
    let canvas = this.renderer.createElement('canvas');
    canvas.width = this.width;
    canvas.height = this.height;

    let context = canvas.getContext('2d') as CanvasRenderingContext2D

    let fontSize = this.width / (2 / (this._fontScale / 100));
    context.font = `${this._fontWeight} ${fontSize}px sans-serif`;

    context.fillStyle = this.bgColor;
    context.fillRect(0, 0, this.width, this.height);

    if (this.showShadow) {
      context.shadowColor = "black";
      context.shadowOffsetX = 0;
      context.shadowOffsetY = 0;
      context.shadowBlur = 5;
    }

    context.textAlign = "center";
    context.fillStyle = this.textColor;
    context.fillText(this.initial, this.width / 2, this.height - (this.height / 2) + (this.fontSize / 3));

    let result = canvas.toDataURL('image/' + this._pictureFormat);
    console.log(result);
    return canvas.toDataURL("image/" + this._pictureFormat);
  }

}
