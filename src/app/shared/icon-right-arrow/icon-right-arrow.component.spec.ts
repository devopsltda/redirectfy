import { ComponentFixture, TestBed } from '@angular/core/testing';

import { IconRightArrowComponent } from './icon-right-arrow.component';

describe('IconRightArrowComponent', () => {
  let component: IconRightArrowComponent;
  let fixture: ComponentFixture<IconRightArrowComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [IconRightArrowComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(IconRightArrowComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
