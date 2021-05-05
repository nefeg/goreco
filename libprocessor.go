package goreco

import (
	"fmt"
	"gocv.io/x/gocv"
)

type processor struct {
	Processor
	process func(detectionResult gocv.Mat, threshold float32, all bool) []Box
}

func (p *processor) Process(detectionResult gocv.Mat, threshold float32, all bool) []Box {
	return p.process(detectionResult, threshold, all)
}

func NewProcessorSSD() Processor {
	p := &processor{}
	p.process = postProcessSSD
	return p
}

func NewProcessorYOLO() Processor {
	p := &processor{}
	p.process = postProcessYOLO
	return p
}

//func (p *processor)Process(detectionResult gocv.Mat, threshold float32 ) []Box{
//
//}

// performDetection analyzes the results from the detector network,
// which produces an output blob with a shape 1x1xNx7
// where N is the number of detections, and each detection
// is a vector of float values
// [batchId, classId, confidence, left, top, right, bottom]
func postProcessSSD(results gocv.Mat, threshold float32, all bool) (b []Box) {

	for i := 0; i < results.Total(); i += 7 {

		confidence := results.GetFloatAt(0, i+2)
		classId := int(results.GetFloatAt(0, i+1))

		//fmt.Println(results.Cols(), results.Size())
		//fmt.Println(results.GetFloatAt(0, i+1), results.GetFloatAt(0, i+2), results.GetFloatAt(0, i+3), results.GetFloatAt(0, i+4))
		//fmt.Println(results.GetFloatAt(0, i+5), results.GetFloatAt(0, i+6), results.GetFloatAt(0, i+7), results.GetFloatAt(0, i+8))
		//fmt.Println(results.GetFloatAt(0, i+9), results.GetFloatAt(0, i+10), results.GetFloatAt(0, i+11), results.GetFloatAt(0, i+12))

		//if confidence > threshold {
		if confidence > threshold {
			if !all && classId != 1 {
				continue
			}
			box := Box{
				class: classId,
				rect: struct {
					min [2]float32
					max [2]float32
				}{
					[2]float32{results.GetFloatAt(0, i+3), results.GetFloatAt(0, i+4)},
					[2]float32{results.GetFloatAt(0, i+5), results.GetFloatAt(0, i+6)},
				},
				conf: confidence,
			}
			//fmt.Println(box)
			b = append(b, box)
		}
	}
	return b
}

/**
The outputs object are vectors of lenght 85

4x the bounding box (centerx, centery, width, height)
1x box confidence
80x class confidence
*/
func postProcessYOLO(results gocv.Mat, threshold float32, all bool) (b []Box) {

	fts, _ := results.DataPtrFloat32()

	fmt.Println(len(fts))

	fmt.Println(results.Size(), results.Total())
	fmt.Println(results.GetFloatAt(0, 0), results.GetFloatAt(0, 1), results.GetFloatAt(0, 2), results.GetFloatAt(0, 3))
	fmt.Println(results.GetFloatAt(0, 4), results.GetFloatAt(0, 5))
	threshold = 0.2

	for i := 0; i < results.Total(); i += 85 {
		tensor := fts[i : i+85]

		for i, v := range tensor[5:] {
			if v != 0 {
				fmt.Println(i, v)
			}
		}
		//boxConfidence := tensor[4]
		classConfidence, classId := FindMaxFloat32(tensor[5:])

		if classId == 0 {
			//if classId ==0 ||classConfidence <threshold{
			continue
		}

		box := Box{
			class: classId,
			rect: struct {
				min [2]float32
				max [2]float32
			}{
				[2]float32{tensor[0] - tensor[2]/2, tensor[1] - tensor[3]/2},
				[2]float32{tensor[0] + tensor[2]/2, tensor[1] + tensor[3]/2},
			},
			conf: classConfidence,
		}

		fmt.Println(box)
		b = append(b, box)

		//fmt.Println( boxConfidence, box, classId, classConf )

		//confidence := results.GetFloatAt(0, i+2)
		//classId := int(results.GetFloatAt(0, i+1))
		//
		//fmt.Println(confidence, classId)

		//if confidence > threshold {
		//	//if confidence > threshold && classId ==1 {
		//	box := Box{
		//		class: classId,
		//		rect: struct{
		//			min	[2]float32
		//			max [2]float32
		//		}{
		//			[2]float32{results.GetFloatAt(0, i+3), results.GetFloatAt(0, i+4)},
		//			[2]float32{results.GetFloatAt(0, i+5), results.GetFloatAt(0, i+6)},
		//		},
		//		conf: confidence,
		//	}
		//	fmt.Println(box)
		//	b =  append( b, box )
		//}
	}
	return b
}
