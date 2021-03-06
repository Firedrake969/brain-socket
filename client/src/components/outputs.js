import React from 'react';
import { Component } from 'react';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';

class Outputs extends Component {
  constructor(props) {
    super(props);
    this.state = {};
    this.renderOutputs = this.renderOutputs.bind(this);
  }
  renderOutputs() {
    let bases = Object.keys(this.props.outputs).sort();
    return (
      bases.map(name => {
        return (
          <div key={`output-${name}`}>
            <div>{name}</div>
            <div>
              <div>
                {this.props.outputs[name].map((val, idx) => {
                  return (
                    <div key={`output-${name}-${idx}`}>{val}</div>
                  );
                })}
              </div>
            </div>
          </div>
        );
      })
    );
    /*
    <div>
      {this.props.outputs[name].left.map((val, idx) => {
        return (
          <div key={`output-${name}-left-${idx}`}>{val}</div>
        );
      })}
    </div>
    <br />
    <div>
      {this.props.outputs[name].right.map((val, idx) => {
        return (
          <div key={`output-${name}-right-${idx}`}>{val}</div>
        );
      })}
    </div>
    */
  }
  render() {
    return (
      <div>
        {this.renderOutputs()}
      </div>
    );
  }
}

function mapStateToProps(state) {
  return {
    outputs: state.outputs
  };
}

function mapDispatchToProps(dispatch) {
  return bindActionCreators({  }, dispatch);
}

export default connect(mapStateToProps, mapDispatchToProps)(Outputs);