import { ComponentFixture, TestBed } from '@angular/core/testing';

import { McPickBoxComponent } from './mc-pick-box.component';

describe('McPickBoxComponent', () => {
  let component: McPickBoxComponent;
  let fixture: ComponentFixture<McPickBoxComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ McPickBoxComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(McPickBoxComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
