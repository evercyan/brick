const solution = require('./solution.js')

let cases = [
    {
        "inputs": [
            "input"
        ],
        "expects": [
           	"output"
        ],
    }
];

cases.forEach(function(item, i) {
    test('test-' + i, () => {
		let ret = solution(item['inputs'][0])
		let expected = item['expects'][0]
        expect(ret).toBe(expected)
    })
});