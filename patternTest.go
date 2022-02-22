package main

import (
	"fmt"

	"./pattern"
)

func main() {
	/*
				facade := pattern.Facade{}
				facade.MakeOrder("Dd", "dd")

				fmt.Println("")



		biiiiiiiiiiiiiiiiiiiiiiiiiiiilddddderrrrrrrrrr
	*/
	fb := pattern.NewFirstBuilder()

	director := pattern.NewDirector(fb)
	director.BuildHouse()

	fmt.Println()

	sb := pattern.NewSecondBuild()

	director.SetBulder(sb)
	director.BuildHouse()

	/*
								viiiiiiiiiiiiiiiiisssssssssssssssssiiiiiiiiiiiiiiiiiiiiit

									c := &pattern.Car{10, 10}
									b := &pattern.Motorbike{7}

									travel := &pattern.Travel{}

									c.Accept(travel)
									b.Accept(travel)

									sum := &pattern.SumGas{}

									c.Accept(sum)
									b.Accept(sum)


							// cooooooommmmmmmmmmmmmmmaaaandddd

							rec := &pattern.Recipient{}

							oneCommand := &pattern.OneCommand{rec}
							oneCommand.Execute()
							twoCommand := &pattern.TwoCommnad{rec}
							twoCommand.Execute()



						///ccchhhaaan ooooff reeeessp

						hThree := &pattern.HandlerThree{}

						hTwo := &pattern.HandlerTwo{}
						hTwo.SetNext(hThree)

						hOne := &pattern.HandlerOne{}
						hOne.SetNext(hTwo)

						client := &pattern.Client{Name: "Vasya"}
						hOne.Execute(client)

						/// ffactttooryyy mmmmmmmmmmmmmmmmmmeeetod

					pn := pattern.NewPen()
					pl := pattern.NewPencil()

					f(pn)
					f(pl)

				}
				func f(iD pattern.IDeviceWrite) {
					fmt.Println(iD)
					iD.UseDevice()
				}






			strategy1 := &pattern.FirstStrategy{}
			testObject := pattern.InitStrategy(strategy1)
			testObject.UseStrategy()

			strategy2 := &pattern.SecondStrategy{}
			testObject.SetStrategy(strategy2)
			testObject.UseStrategy()

			strategy3 := &pattern.ThirdStrategy{}
			testObject.SetStrategy(strategy3)
			testObject.UseStrategy()


		context := pattern.NewContext()

		context.FirstOperation()
		context.SecondOperation()
		context.ThirdOperation()

		context.FirstOperation()
		context.SecondOperation()
	*/

}
